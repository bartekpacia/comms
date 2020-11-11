package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jacobsa/go-serial/serial"
)

var (
	options = serial.OpenOptions{
		PortName:        "/dev/tty.usbserial-141210",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 1,
		ParityMode:      serial.PARITY_NONE,
	}
)

func main() {
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
