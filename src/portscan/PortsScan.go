package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s ip-addr portlist", os.Args[0])
		os.Exit(1)
	}
	ip := os.Args[1]
	ports := make(chan string)
	for i := 2; i < len(os.Args); i++ {
		target := ip + ":" + os.Args[i]
		tcpAddr, err := net.ResolveTCPAddr("tcp4", target)
		checkError(err)
		go checkPort(tcpAddr, os.Args[i], ports)
	}
	for {
		select {
		case port := <-ports:
			fmt.Println(port)
		case <-time.After(2 * time.Second):
			return
		}

	}
}

func checkPort(target *net.TCPAddr, port string, ports chan string) {
	conn, err := net.DialTCP("tcp", nil, target)
	if err == nil {
		conn.Close()
		ports <- port
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
