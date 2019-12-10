package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var (
	ip = flag.String("ip", "127.0.0.1", "server IP")
	connections = flag.Int("conn", 1, "number of socket connections")
)

func main() {
	flag.Usage = func() {
		io.WriteString(os.Stderr, `sockets client generator Example usage: ./client -ip=172.17.0.1 -conn=10`)
		flag.PrintDefaults()
	}
	flag.Parse()

	log.Printf("Connecting to %s", *ip + ":8000")
	var conns []net.Conn
	for i := 0; i < *connections; i++ {
		c, err := net.Dial("tcp", *ip + ":8000")
		if err != nil {
			fmt.Println("Failed to connect", i, err)
			break
		}
		conns = append(conns, c)
		go func(i int, c net.Conn) {
			for {
				if _, err := c.Write([]byte(fmt.Sprintf("hello from %d\n", i))); err != nil {
					fmt.Println("failed to sending data: ", err)
					break
				}
				time.Sleep(time.Second)
			}
		}(i, c)
	}

	log.Printf("Finished initializing %d connections", len(conns))
	tts := time.Second
	if *connections > 100 {
		tts = time.Microsecond * 5
	}
	for {
		for i := 0; i < len(conns); i++ {
			time.Sleep(tts)
			conn := conns[i]
			log.Printf("Conn %d sending message", i)
			reader := bufio.NewReader(conn)
			if msg, err := reader.ReadString('\n'); err != nil {
				fmt.Printf("Failed to receive pong: %v\n", err)
			}else{
				fmt.Printf("receive msg from %v: %s\n", i , msg)
				conn.Write([]byte(msg))
			}
		}
	}
}



