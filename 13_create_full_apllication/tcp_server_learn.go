package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start the server...")
	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("error listening", err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting", err.Error())
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("received data: %v", string(buf[:len]))
	}
}
