package main

import (
	"fmt"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

func main() {
	options := serial.OpenOptions{
		PortName:        "/dev/tty.usbserial-141230",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 8,
		ParityMode:      serial.PARITY_NONE,
	}

	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("error opening serial port: %v\n", err)
	}

	defer port.Close()

	for {
		b := make([]byte, 8)
		n, err := port.Read(b)
		if err != nil {
			log.Fatalf("error reading from serial port: %v\n", err)
		}

		fmt.Printf("read %d bytes: %d\n", n, b[0])
	}
}
