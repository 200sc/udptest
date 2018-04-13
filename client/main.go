package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Client starting")
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	if err != nil {
		fmt.Println(err)
		return
	}

	Conn, err := net.DialUDP("udp", nil, ServerAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer Conn.Close()
	i := 0
	fmt.Println("Client sending messages")
	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
}
