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
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("invalid arguments. specify address of the server")
	}
	if err := Connect(args[0]); err != nil {
		log.Fatalf("cannot connect to %s; %v", args[0], err)
	}
}

func Connect(addr string) error {
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
