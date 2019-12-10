package main

import (
	"bufio"
	"context"
	"eagle/utils"
	"eagle/utils/network"
	"flag"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	address       *string
	pprof_address = "0.0.0.0:6060"
	epoller       *network.EPoll
)

func init() {
}

func Parse() {
	address = flag.String("server", "0.0.0.0:8000", "server address")
	flag.Parse()
}

func main() {
	Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	setLimit()

	go func() {
		if err := http.ListenAndServe(pprof_address, nil); err != nil {
			utils.Logger.Fatalf("pprof failed: %v", err)
		}
	}()

	netListener, err := net.Listen("tcp", *address)
	if err != nil {
		utils.Logger.Fatalf("failed to listen: %v", err)
	}
	utils.Logger.Println("listen on: ", *address)

	epoller, err = network.MkEpoll()
	if err != nil {
		panic(err)
	}

	go Start(ctx)

	go handleConnection(netListener, ctx)

	gracefulExit(cancel)
	<-ctx.Done()
}

func setLimit() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
}

func handleConnection(netListener net.Listener, ctx context.Context) {
	ctx2, cancel2 := context.WithCancel(ctx)
	defer cancel2()
	for {
		select {
		case <- ctx2.Done():
			err := netListener.Close()
			if err != nil {
				log.Println("shutdown server error ", err)
			}
		default:
			conn, err := netListener.Accept()
			if err != nil {
				continue
			}

			if err := epoller.Add(conn); err != nil {
				utils.Logger.Printf("Failed to add connection %v", err)
				conn.Close()
			}
		}
	}
}

func Start(ctx context.Context) {
	ctx2, cancel2 := context.WithCancel(ctx)
	defer cancel2()
	for {
		connections, err := epoller.Wait()
		if err != nil {
			utils.Logger.Printf("Failed to epoll wait %v", err)
			continue
		}

		select {
		case <-ctx2.Done():
			for _, conn := range connections {
				if err := epoller.Remove(conn); err != nil {
					utils.Logger.Printf("Failed to remove %v", err)
				}
				conn.Close()
			}
			utils.Logger.Printf("remove all connections")
			return
		default:
			for _, conn := range connections {
				if conn == nil {
					break
				}
				reader := bufio.NewReader(conn)
				if msg, err := reader.ReadString('\n'); err != nil {
					if err := epoller.Remove(conn); err != nil {
						utils.Logger.Printf("Failed to remove %v", err)
					}
				} else {
					// process data
					conn.Write([]byte(msg))
					utils.Logger.Printf("%s", msg)
				}
			}
		}
	}
}

func gracefulExit(cancelFunc context.CancelFunc) {
	now := time.Now()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch
	log.Println("got a signal ", sig)
	log.Println("--------------- start cleaning ---------------")

	cancelFunc()
	uptime := time.Since(now)
	log.Println("          shutdown server success.            ")



	log.Println("           uptime: ", uptime, "               ")
	log.Println("------------------ exited --------------------")
}
