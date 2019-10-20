package genericWifiLed

import (
	"fmt"
	"testing"
)

const TEST_UNIT_IP = "192.168.1.137"

func Test_genericWifiLed_DimTo_offB(t *testing.T) {
	wifiController := NewWifiLedController(TEST_UNIT_IP, 0)
	wifiController.DimTo(0,0,0,0,0)
}

func Test_genericWifiLed_DimTo_fullRGB(t *testing.T) {
	wifiController := NewWifiLedController(TEST_UNIT_IP, 0)
	wifiController.DimTo(255,255,255,0,0)
}

func Test_genericWifiLed_generateChecksum_blank(t *testing.T) {
	actual := addChecksum([]byte{})
	if len(actual) != 0 {
		t.Error("Expected to get back empty set but got", actual)
	}
}

func Test_genericWifiLed_generateChecksum_blue7percent(t *testing.T) {
	expected := []byte {0x31, 0x00, 0x00, 0x07, 0x0, 0x0, 0xF0, 0x0F, 0x37}
	input := []byte {0x31, 0x00, 0x00, 0x07, 0x0, 0x0, 0xF0, 0x0F}
	actual := addChecksum(input)

	if len(expected) != len(actual){
		t.Error("invalid")
	}

	fmt.Println("ACTUAL  : ", actual)
	fmt.Println("EXPECTED: ", expected)

}