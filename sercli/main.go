package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Sercli starting")

	lst, err := listenConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	bcst, err := broadcastConn()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer lst.Close()
	defer bcst.Close()
	i := 0
	fmt.Println("Client sending messages")
	go func() {
		for {
			msg := strconv.Itoa(i)
			fmt.Println("Sending", i)
			i++
			buf := []byte(msg)
			_, err := bcst.Write(buf)
			if err != nil {
				fmt.Println(msg, err)
			}
			time.Sleep(time.Second * 20)
		}
	}()
	buf := make([]byte, 1024)
	for {
		n, addr, err := lst.ReadFromUDP(buf)
		// Todo: check we didn't send this?
		// UDP shouldn't send to us, though
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

func listenConn() (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp", ":10001")
	if err != nil {
		return nil, err
	}

	cnn, err := net.ListenUDP("udp", addr)
	return cnn, err
}

func broadcastConn() (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp", "255.255.255.255:10001")
	if err != nil {
		return nil, err
	}

	laddr, err := net.ResolveUDPAddr("udp", ":10002")
	if err != nil {
		return nil, err
	}

	return net.DialUDP("udp", laddr, addr)
}
