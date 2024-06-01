package main

import (
	"fmt"
	"net"
	"os"
	"log"
)

// this code is the same as the other server

func main() {
	// Check if a port argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <port>")
		return
	}

	// Get the port from the command-line argument
	port := os.Args[1]

	// Listen for incoming connections on the specified port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()
	fmt.Printf("Server listening on port %s\n", port)

	// setting up logs
	logFile, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Unable to open log file:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	for {
		// Accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a separate goroutine
		go handleConnection(conn)
	}
}
