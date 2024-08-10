package genericWifiLed

import (
	"fmt"
	"testing"
)

const TEST_UNIT_IP = "192.168.1.120" // set this to your local device to test against
const TEST_UNIT_PORT = "5577"
const TEST_TIMEOUT = 5

func Test_genericWifiLed_DimTo_offB(t *testing.T) {
	wifiController := NewController(TEST_UNIT_IP, TEST_UNIT_PORT, TEST_TIMEOUT)
	wifiController.DimTo(0, 0, 0, 0, 0)
	t.Log("The rgb strip should now be off")
}

func Test_genericWifiLed_DimTo_fullRGB(t *testing.T) {
	wifiController := NewController(TEST_UNIT_IP, TEST_UNIT_PORT, TEST_TIMEOUT)
	wifiController.DimTo(255, 255, 255, 0, 0)
}

func Test_genericWifiLed_DimTo_AboveRangeRed(t *testing.T) {
	wifiController := NewController(TEST_UNIT_IP, TEST_UNIT_PORT, TEST_TIMEOUT)
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
		t.Error("invalid byte length")
	}

	fmt.Println("ACTUAL  : ", actual)
	fmt.Println("EXPECTED: ", expected)

}

func Test_genericWifiLed_multipleIp(t *testing.T) {
	firstIp := "192.168.1.2"
	secondIp := "192.168.1.3"
	input := firstIp + "," + secondIp
	wifiController := NewController(input, "", 0)

	checkArray(t, wifiController.ipAddresses, []string{firstIp, secondIp})
}

func checkArray(t *testing.T, actual []string, expected []string) {
	if len(actual) != len(expected) {
		t.Error("Expected to have", len(expected), "rows but got", len(actual))
	}

	for i := range actual {
		checkString(t, actual[i], expected[i])
	}
}

func checkString(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Error("EXPECTED : ", expected, "\nACTUAL   :", actual)
	}
}
