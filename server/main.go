package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
)

var port = flag.String("p", "8080", "define the address of the port.")
var users map[net.Conn]bool

func main() {
	flag.Parse()
	p := *port

	users = make(map[net.Conn]bool)

	if err := listen(p); err != nil {
		log.Fatal(err)
	}
}

// listen opens a TCP connection and waits for incoming client connections.
func listen(port string) error {
	if len(port) == 0 || strings.TrimSpace(port) == "" {
		return fmt.Errorf("invalid port %s", port)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return fmt.Errorf("cannot listen tcp on port=%s; %v", port, err)
	}

	log.Printf("listening on: %s\n", listener.Addr().String())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		log.Printf("%s connected\n", conn.RemoteAddr().String())
		users[conn] = true

		go func(conn net.Conn) {
			scanner := bufio.NewScanner(conn)
			defer conn.Close()
			for scanner.Scan() {
				msg := scanner.Text()
				for user := range users {
					if user.RemoteAddr().String() != conn.RemoteAddr().String() {
						fmt.Fprintf(user, "%s - %s", conn.RemoteAddr().String(), msg)
					}
				}
			}
			delete(users, conn)
			log.Printf("%s disconnected\n", conn.RemoteAddr().String())
		}(conn)
	}
}
