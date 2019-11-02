package genericWifiLed

import (
	"fmt"
	"testing"
)

const TEST_UNIT_IP = "192.168.1.120" // set this to your local device to test against

func Test_genericWifiLed_DimTo_offB(t *testing.T) {
	wifiController := NewController(TEST_UNIT_IP, "")
	wifiController.DimTo(0, 0, 0, 0, 0)
}

func Test_genericWifiLed_DimTo_fullRGB(t *testing.T) {
	wifiController := NewController(TEST_UNIT_IP, "")
	wifiController.DimTo(255, 255, 255, 0, 0)
}

func Test_genericWifiLed_DimTo_AboveRangeRed(t *testing.T) {
	wifiController := NewController(TEST_UNIT_IP, "")
	wifiController.DimTo(600, 0, 0, 0, 0)
}

func Test_genericWifiLed_generateChecksum_blank(t *testing.T) {
	actual := addChecksum([]byte{})
	if actual[0] != 0x0 {
		t.Error("Expected to get back empty set but got", actual)
	}
}

func Test_genericWifiLed_generateChecksum_blue7percent(t *testing.T) {
	expected := []byte{0x31, 0x00, 0x00, 0x07, 0x0, 0x0, 0xF0, 0x0F, 0x37}
	input := []byte{0x31, 0x00, 0x00, 0x07, 0x0, 0x0, 0xF0, 0x0F}
	actual := addChecksum(input)

	if len(expected) != len(actual) {
		t.Error("invalid")
	}

	fmt.Println("ACTUAL  : ", actual)
	fmt.Println("EXPECTED: ", expected)

}
