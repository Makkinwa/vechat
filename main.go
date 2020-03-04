package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	start()
}

func start() {
	// Listen some address
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	} else {
		for {
			// Listen for an incoming connection.
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				os.Exit(1)
			}
			// Handle connections in a new goroutine.

			go handleRequest(conn)
		}
	}
}

func handleRequest(conn net.Conn ) {
	buf := make([]byte, 255)
	print("New connection")
	// Read the incoming connection into the buffer.
	_ , err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	} else {
		print(string(buf))
		mysql_handle(buf)
	}
}
