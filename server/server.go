package main

import (
	"bufio"
	"eagle/utils"
	"eagle/utils/network"
	"flag"
	"net"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"syscall"
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

	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

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

	go Start()

	for {
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

func Start() {
	for {
		connections, err := epoller.Wait()
		if err != nil {
			utils.Logger.Printf("Failed to epoll wait %v", err)
			continue
		}

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
