package main

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
	"os"
)

//func send() error {
//	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
//	if err != nil {
//		return errors.WithStack(err)
//	}
//
//	serverAddr
//}

func recv() error {
	serverAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:8080")
	if err != nil {
		return errors.WithStack(err)
	}

	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		return errors.WithStack(err)
	}
	defer conn.Close()

	for {
		println("reading...")
		var clientAddr *net.UDPAddr
		buffer := make([]byte, 1024)
		_, clientAddr, err = conn.ReadFromUDP(buffer)
		if err != nil {
			return errors.WithStack(err)
		}

		fmt.Printf("recv: %s\n---\n%s---\n", clientAddr, buffer)

		_, err = conn.WriteTo(buffer, clientAddr)
		if err != nil {
			return errors.WithStack(err)
		}
	}
}

func main() {
	pid := os.Getpid()
	println("pid:", pid)

	err := recv()
	if err != nil {
		panic(err)
	}
}
