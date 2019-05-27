package main

/*
Syntax Go. Homework 7.3 - goroutines and chanels
Task: Correct timeserver shutdown
Nick Nikulin, dated Мay 25, 2019
Teacher, Sergey Iryupin.
*/

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	go waitExit()
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
		handleConn(conn)
	}
}

func waitExit() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			os.Exit(0)
		}
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
