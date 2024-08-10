package genericWifiLed

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const COMMANDGROUP_SETCOLOR = 0x31
const COMMANDGROUP_SETMODE = 0x61
const COMMANDGROUP_SETPOWER = 0x71
const SOCKET_TRUE = 0xF0
const SOCKET_FALSE = 0x0F
const SOCKET_ON = 0x23
const SOCKET_OFF = 0x24

type WifiLedController struct {
	ipAddresses []string
	port        int
	timeout     int
}

func NewController(_ipAddress string, port string, timeout int) (output WifiLedController) {
	output.ipAddresses = strings.Split(_ipAddress, ",")
	output.port, _ = strconv.Atoi(port)
	if output.port == 0 {
		output.port = 5577
	}
	output.timeout = timeout
	if output.timeout < 1 {
		output.timeout = 1
	}
	return
}

func (w *WifiLedController) DimTo(red int, green int, blue int, warmWhite int, coolwhite int) {
	payload := generateDimCommand(red, green, blue, warmWhite, coolwhite)
	w.dialAndSendThenClose(payload)
}
func (w *WifiLedController) On() {
	payload := generateOnCommand()
	w.dialAndSendThenClose(payload)
}
func (w *WifiLedController) Off() {
	payload := generateOffCommand()
	w.dialAndSendThenClose(payload)
}

func generateOnCommand() []byte {
	payload := []byte{COMMANDGROUP_SETPOWER, SOCKET_ON, SOCKET_FALSE}
	return addChecksum(payload)
}
func generateOffCommand() []byte {
	payload := []byte{COMMANDGROUP_SETPOWER, SOCKET_OFF, SOCKET_FALSE}
	return addChecksum(payload)
}
func generateDimCommand(red int, green int, blue int, warmWhite int, coolWhite int) []byte {
	payload := []byte{COMMANDGROUP_SETCOLOR, byte(red), byte(green), byte(blue), byte(warmWhite), byte(coolWhite), SOCKET_TRUE, SOCKET_FALSE}
	return addChecksum(payload)
}

func addChecksum(input []byte) []byte {
	checksum := byte(0)
	for i := range input {
		checksum += input[i]
	}
	return append(input, checksum)
}

func (w *WifiLedController) dialAndSendThenClose(payload []byte) (err error) {
	timeout := time.Duration(w.timeout) * time.Second
	for _, ip := range w.ipAddresses {
		fmt.Printf("  LED => %s ", ip)
		connection, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(w.port), timeout)
		if err != nil {
			fmt.Println(err)
			return err
		}

		_, err = connection.Write(payload)
		if err != nil {
			fmt.Println(err)
		}

		err = connection.Close()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("+\n")
	}
	return
}
