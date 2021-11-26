package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type outWriter struct{}

func (o *outWriter) Write(p []byte) (n int, err error) {
	n, err = fmt.Fprintf(os.Stdout, "%s", string(p))
	if err != nil {
		return n, err
	}

	fmt.Println()
	return n, err
}

func main() {
	addr := ":8080"
	args := os.Args[1:]
	if len(args) != 0 {
		log.Printf("%s will be used as the server address.", args[0])
		addr = args[0]
	}
	if err := Connect(addr); err != nil {
		log.Fatalf("cannot connect to %s; %v", addr, err)
	}
}

func Connect(addr string) error {
	log.Println("connecting to ", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	go func(conn net.Conn) {
		ow := &outWriter{}
		if _, err := io.Copy(ow, conn); err != nil {
			log.Fatal(err)
		}
	}(conn)

	// Send messages from os.Stdin to conn
	if _, err := io.Copy(conn, os.Stdin); err != nil {
		log.Fatal(err)
	}
	return nil
}
