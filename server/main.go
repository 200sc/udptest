package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server starting")
	ServerAddr, err := net.ResolveUDPAddr("udp", ":10001")
	if err != nil {
		fmt.Println(err)
		return
	}

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ServerConn.Close()

	buf := make([]byte, 1024)
	fmt.Println("Server waiting")
	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
