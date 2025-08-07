package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "45.236.129.53:7171")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Simulamos handshake tipo Tibia 8.6
	packet := []byte{
		0x0A, 0x00, // Packet size (little endian): 10 bytes
		0x0A,       // OS = 10 (Windows)
		0x00,       // client version high byte
		0x5C, 0x03, // client version low byte = 860 (0x035C)
		0x00, 0x00, 0x00, 0x00, // dat, spr, pic signatures (vacios)
	}
	conn.Write(packet)

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Respuesta del servidor (%d bytes):\n% X\n", n, buffer[:n])
}
