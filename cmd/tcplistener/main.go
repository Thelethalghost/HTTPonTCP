package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Thelethalghost/httpfromtcp/internal/request"
)

func main() {

	listener, err := net.Listen("tcp", ":42069")

	if err != nil {
		log.Fatal("error", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		request, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatal(err)
		}

		method := request.RequestLine.Method
		target := request.RequestLine.RequestTarget
		version := request.RequestLine.HttpVersion

		fmt.Printf("Request Line: \n - Method: %s\n - Target: %s\n - Version: %s\n", method, target, version)
	}
}
