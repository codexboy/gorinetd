package main

import (
	"flag"
	"fmt"
	"net"
)

func ConnRW(conn net.Conn) {
	defer conn.Close()
	fmt.Println(conn.RemoteAddr().String())
	buftmp := make([]byte,1024*10)
	for true {
		rc,err := conn.Read(buftmp)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(string(buftmp[:rc]))
	}
}

func main() {
	host := flag.String("h", "127.0.0.1", "host address")
	port := flag.Int("p", 9999, "host port")

	flag.Parse()

	fmt.Println("Used:", fmt.Sprintf("%s:%d", *host, *port))
	sock, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sock.Close()
	for true {
		conn, err := sock.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go ConnRW(conn)
	}
}
