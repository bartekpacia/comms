package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

var (
	portName   string
	baudRate   uint
	dataBits   uint
	stopBits   uint
	parityMode int
)

func init() {
	Bootstrap()
}

func main() {
	flag.Parse()

	options := serial.OpenOptions{
		PortName:        portName,
		BaudRate:        baudRate,
		DataBits:        dataBits,
		StopBits:        stopBits,
		MinimumReadSize: 1,
		ParityMode:      serial.ParityMode(parityMode),
	}

	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("error opening serial port: %v\n", err)
	}
	defer port.Close()

	for {
		b := make([]byte, 1)
		n, err := port.Read(b)
		if err != nil {
			log.Fatalf("error reading from serial port: %v\n", err)
		}

		fmt.Println("bytes read:", n, "value:", b[0])
	}
}
