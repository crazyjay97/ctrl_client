package core

import (
	"bufio"
	"com.lierda.wsn.vc/util"
	"go.bug.st/serial"
	"log"
)

func ReadSerialList() []string {
	portList, _ := serial.GetPortsList()
	return portList
}

var Msg = make(chan string, 10)

var flag = true

func ClosePort() {
	flag = false
}

var Write func(str string) = nil

func OpenPort(port string) error {
	ser, err := serial.Open(port, &serial.Mode{BaudRate: 115200})
	Write = func(str string) {
		n, err := ser.Write([]byte(str + "\n\r"))
		log.Println(err)
		Msg <- util.CurrentTimeString() + " S: " + str
		log.Printf("Sent %v bytes\n", n)
	}

	if err != nil {
		return err
	}
	flag = true
	reader := bufio.NewReader(ser)
	go func() {
		for {
			if flag {
				line, _, err := reader.ReadLine()
				if err != nil {
					log.Fatal(err)
				}
				Msg <- util.CurrentTimeString() + " R: " + string(line)
			} else {
				ser.Close()
				Write = nil
				break
			}
		}
	}()
	return nil
}
