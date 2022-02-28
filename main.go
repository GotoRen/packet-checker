package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"test/internal"
)

func main() {
	intf, err := net.InterfaceByName(os.Getenv("INTERFACE"))
	if err != nil {
		fmt.Println("Failed to read .env", err)
	}

	recv4sock, err := internal.RecvIPv4RawSocket(intf)
	if err != nil {
		fmt.Println("Failed to generate receive for IPv4 descriptor", err)
	}

	for {
		buf := make([]byte, 1500)
		size, _, err := syscall.Recvfrom(recv4sock, buf, 0)
		if err != nil {
			fmt.Println("Failed to read packet", err)
		}

		internal.PrintPacketInfo(buf[:size])
	}
}
