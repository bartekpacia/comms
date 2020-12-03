package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

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

func convert(b byte) byte {
	shift := b - 48
	var mask byte
	mask = 0b00000001 << shift

	if shift > 7 {
		fmt.Printf("uart_leds: convert(%b): shift > 7, something is wrong\n", b)
	}

	return mask
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

	fmt.Println("uart_leds: toggle LEDs on and off")
	fmt.Println("uart_leds: enter 8 to exit")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("enter a single byte: ")
		b, err := reader.ReadByte()
		if err != nil {
			log.Fatalln("error reading from stdin:", err)
		}

		if b == '\n' {
			break
		}

		mask := convert(b)

		n, err := port.Write([]byte{mask})
		if err != nil {
			log.Fatalln("error writing to serial port:", err)
		}

		fmt.Printf("wrote %d bytes to serial port (%d)\n", n, mask)
	}

	fmt.Println("finished :)")
}

// TODO: Move to a goroutine
// output := make([]byte, 1)
// n, err = port.Read(output)
// if err != nil {
// 	log.Fatalln("error reading from serial port:", err)
// }
// b = output[0]
// fmt.Printf("read %d bytes from serial port (%d)\n", n, b)
