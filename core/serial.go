package core

import (
	"bufio"
	"go.bug.st/serial"
	"log"
)

func ReadSerialList() []string {
	portList, _ := serial.GetPortsList()
	for _, port := range portList {
		log.Println(port)
	}
	return portList
}

var Msg = make(chan string, 10)

func OpenPort(port string) error {
	ser, err := serial.Open(port, &serial.Mode{BaudRate: 115200})
	if err != nil {
		return err
	}
	reader := bufio.NewReader(ser)
	go func() {
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			Msg <- string(line)
		}
	}()
	return nil
}
