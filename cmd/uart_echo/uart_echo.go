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

	fmt.Print("uart_echo: a tiny program that sends a byte to ")
	fmt.Printf("microcontroller and then receives ")
	fmt.Printf("the very same byte.\n")
	fmt.Println("uart_echo: enter -1 to stop")

	for {
		var value int
		fmt.Print("enter a single byte to be sent (integer, 0-255): ")
		_, err := fmt.Scanf("%d", &value)
		if err != nil {
			log.Fatalln("error reading from stdin:", err)
		}

		if value == -1 {
			break
		}

		if value < 0 || value > 255 {
			fmt.Printf("uart_echo: error: %d overflows byte\n", value)
			break
		}

		inputByte := byte(value)
		fmt.Printf("\"%b\" will be sent\n", inputByte)
		n, err := port.Write([]byte{inputByte})
		if err != nil {
			log.Fatalln("error writing to serial port:", err)
		}
		fmt.Printf("wrote %d bytes (\"%d\") to serial port\n", n, inputByte)

		output := make([]byte, 1)
		n, err = port.Read(output)
		if err != nil {
			log.Fatalln("error reading from serial port:", err)
		}
		outputByte := output[0]

		fmt.Printf("read %d bytes (\"%d\") from serial port \n", n, outputByte)
	}

	fmt.Println("finished :)")
}
