package modbustest

import (
	"fmt"
	"time"
	"log"
	"github.com/goburrow/modbus"
)

func Modbustest() {
	serv := mbserver.NewServer()
	err := serv.ListenTCP("127.0.0.1:1502")
	if err != nil {
		log.Printf("%v\n", err)
	}

	err := serv.ListenTCP("0.0.0.0:3502")
	if err != nil {
		log.Printf("%v\n", err)
	}

	err := s.ListenRTU(&serial.Config{
		Address:  "/dev/ttyUSB0",
		BaudRate: 115200,
		DataBits: 8,
		StopBits: 1,
		Parity:   "N",
		Timeout:  10 * time.Second})
	if err != nil {
		t.Fatalf("failed to listen, got %v\n", err)
	}

	err := s.ListenRTU(&serial.Config{
		Address:  "/dev/ttyACM0",
		BaudRate: 9600,
		DataBits: 8,
		StopBits: 1,
		Parity:   "N",
		Timeout:  10 * time.Second,
		RS485: serial.RS485Config{
			Enabled: true,
			DelayRtsBeforeSend: 2 * time.Millisecond,
			DelayRtsAfterSend: 3 * time.Millisecond,
			RtsHighDuringSend: false,
			RtsHighAfterSend: false,
			RxDuringTx: false,
			}}
		)
	if err != nil {
		fmt.Fatalf("failed to listen, got %v\n", err)
	}

	defer serv.Close()
}
