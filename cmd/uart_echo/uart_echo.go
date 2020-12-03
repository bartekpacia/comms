package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

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
	flag.StringVar(&portName, "port", "/dev/tty.*", "port to listen on")
	flag.UintVar(&baudRate, "baud", 9600, "baud rate in bits per second")
	flag.UintVar(&dataBits, "dbits", 8, "the number of data bits in a single frame")
	flag.UintVar(&stopBits, "sbits", 1, "the number of stop bits in a single frame")
	flag.IntVar(&parityMode, "pmode", 1, "parity mode, none = 0, odd = 1, even = 2")
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
		log.Fatalln("error opening serial port:", err)
	}
	defer port.Close()

	fmt.Println("uart_echo: a tiny program that sends a byte to")
	fmt.Println("uart_echo: a microcontroller and then receives")
	fmt.Println("uart_echo: the very same byte.")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("enter a single byte: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("error reading from stdin:", err)
		}

		if strings.TrimSpace(text) == "STOP" {
			break
		}

		b := []byte(text)[0]
		n, err := port.Write([]byte{b})
		if err != nil {
			log.Fatalln("error writing to serial port:", err)
		}
		fmt.Printf("wrote %d bytes to serial port (%d)\n", n, b)

		output := make([]byte, 1)
		n, err = port.Read(output)
		if err != nil {
			log.Fatalln("error reading from serial port:", err)
		}
		b = output[0]

		fmt.Printf("read %d bytes from serial port (%d)\n", n, b)
	}

	fmt.Println("finished :)")
}
