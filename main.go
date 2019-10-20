package main

import (
	"fmt"
	"os"
	"strconv"
	"wifiled/lib/genericWifiLed"
)

func main() {
	commandLineArguments := os.Args
	commandLineArgumentsLength := len(commandLineArguments)
	command := ""
	if commandLineArgumentsLength >= 2 {
		command = commandLineArguments[1]
	}

	if command == "" {
		displayHelpText()
		return
	}

	//TODO decide how to set ip and port -- maybe an .env file
	wifiController := genericWifiLed.NewWifiLedController("192.168.1.137", 0)
	if command == "on" {
		// todo use the real on command
		wifiController.DimTo(255, 255,255,255,255)
	}else if command == "off" {
		// todo user the real off command
		wifiController.DimTo(0, 0,0,0,0)
	}else if command == "dim" {
		if commandLineArgumentsLength >= 5 {
			redValue := convertStringToBoundedInt(commandLineArguments[2])
			greenValue := convertStringToBoundedInt(commandLineArguments[3])
			blueValue := convertStringToBoundedInt(commandLineArguments[4])
			wifiController.DimTo(redValue, greenValue, blueValue,0,0)
		}else{
			fmt.Println("invalid RGB parameters")
			displayHelpText()
		}
	}else{
		displayHelpText()
	}
}

func displayHelpText()  {
	fmt.Println("wifiled is a program to send commands to generic wifi led controllers on the network")
	fmt.Println("wifiled on -- send on command")
	fmt.Println("wifiled off -- send off command")
	fmt.Println("wifiled dim <RED> <GREEN> <BLUE> -- send on command")
}

func convertStringToBoundedInt(input string) int {
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	if value > 255 {
		return 255
	}
	if value < 0 {
		return 0
	}
	return value
}