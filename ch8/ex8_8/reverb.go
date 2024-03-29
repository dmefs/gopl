package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer c.Close()

	input := bufio.NewScanner(c)
	response := make(chan struct{})

	go func() {
		for input.Scan() {
			go echo(c, input.Text(), 1*time.Second)
		}
	}()
	ticker := time.NewTicker(10 * time.Second)

	select {
	case <-ticker.C:
		fmt.Println("Close idle connection.")
		return
	case <-response:
		ticker.Reset(10 * time.Second)
	}
}
