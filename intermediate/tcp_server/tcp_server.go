package main

import (
	"log"
	"net"
)

// Solution for the reversing tcp server challenge

// TCPServer runs a tcp server which listens to incoming strings and returns them reversed back to the client
func TCPServer(ready chan bool) {
	// listen for incoming connections
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	// tell the client that the server is ready
	ready <- true
	for {
		// listen for an incoming connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// handle connections in goroutines
		go handleRequest(conn)
	}
}

// handleRequest is handle function to run for each incoming string
func handleRequest(conn net.Conn) {
	// close connection after return
	defer conn.Close()
	// make a buffer to hold incoming data
	buf := make([]byte, maxBufferSize)
	// read the incoming connection into the buffer
	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	// send the response back to client
	result := reverse(string(buf[:n]))
	conn.Write([]byte(result))
}

// reverse returns the given string reversed
func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
