package genericWifiLed

import (
	"fmt"
	"net"
	"strconv"
)

const COMMANDGROUP_SETCOLOR = 0x31
const COMMANDGROUP_SETMODE = 0x61
const COMMANDGROUP_SETPOWER = 0x71
const SOCKET_TRUE = 0xF0
const SOCKET_FALSE = 0x0F
const SOCKET_ON = 0x23
const SOCKET_OFF = 0x24

type WifiLedController struct {
	ipAddress string
	port      int
	connection *net.Conn
}

func NewWifiLedController(_ipAddress string, portIfNotDefault int) (output WifiLedController) {
	output.ipAddress = _ipAddress
	if portIfNotDefault > 0 {
		output.port = portIfNotDefault
	}else{
		output.port = 5577
	}
	return
}

func (w *WifiLedController) DimTo(red int, green int, blue int, warmWhite int, coolwhite int) {
	payload := generateDimCommand(red, green, blue, warmWhite, coolwhite)
	w.dialAndSend(payload)
}

func generateDimCommand(red int, green int, blue int, warmWhite int, coolwhite int) []byte {
	payload := []byte { COMMANDGROUP_SETCOLOR, byte(red), byte(green), byte(blue), byte(warmWhite), byte(coolwhite), SOCKET_TRUE, SOCKET_FALSE }
	return addChecksum(payload)
}

func addChecksum(input []byte) []byte {
	checksum := byte(0)
	for i := range input {
		checksum = checksum + input[i]
	}
	return append(input, checksum)
}

func (w *WifiLedController) dialAndSend(payload []byte) (err error) {
	connection, err := net.Dial("tcp", w.ipAddress + ":" + strconv.Itoa(w.port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()

	_, err = connection.Write(payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}