package main

import (
	"fmt"
	"os"
	"wifiled/lib/genericWifiLed"
	"wifiled/lib/toolbox"
)

const VERSION = "0.1"
const ENV_IP_KEY = "wifiled_ip"
const ENV_PORT_KEY = "wifiled_port"

func main() {
	commandLineArguments := os.Args
	commandLineArgumentsLength := len(commandLineArguments)
	command := ""
	if commandLineArgumentsLength >= 2 {
		command = commandLineArguments[1]
	}

	if command == "" {
		displayHelpText("")
		return
	} else if command == "setup" {
		if commandLineArgumentsLength >= 3 {
			err := os.Setenv(ENV_IP_KEY, commandLineArguments[2])
			fmt.Println(commandLineArguments[2])
			if err != nil {
				fmt.Println(err)
			}
			if commandLineArgumentsLength >= 4 {
				err = os.Setenv(ENV_PORT_KEY, commandLineArguments[3])
				if err != nil {
					fmt.Println(err)
				}
			}
			fmt.Println("Device set.")
			return
		} else {
			displayHelpText("invalid setup ip and port parameters")
		}
	}

	ip := os.Getenv(ENV_IP_KEY)
	if ip == "" {
		displayHelpText("Missing ip environment variable, use setup command or check your env")
		return
	}else{
		fmt.Println("IP", ip)
	}
	port := os.Getenv(ENV_PORT_KEY)

	controller := genericWifiLed.NewController(ip, port)
	if command == "on" {
		controller.On()

	} else if command == "off" {
		controller.Off()

	} else if command == "dim" {
		if commandLineArgumentsLength >= 5 {
			const MAX_VALUE = 255
			const MIN_VALUE = 0
			const DEFAULT_VALUE = 0
			redValue := toolbox.ConvertStringToBoundedInt(commandLineArguments[2], MAX_VALUE, MIN_VALUE, DEFAULT_VALUE)
			greenValue := toolbox.ConvertStringToBoundedInt(commandLineArguments[3], MAX_VALUE, MIN_VALUE, DEFAULT_VALUE)
			blueValue := toolbox.ConvertStringToBoundedInt(commandLineArguments[4], MAX_VALUE, MIN_VALUE, DEFAULT_VALUE)
			controller.DimTo(redValue, greenValue, blueValue, 0, 0)
		} else {
			displayHelpText("invalid RGB parameters")
		}

	} else {
		displayHelpText("unknown command")
	}
}

func displayHelpText(errorText string) {
	if errorText != "" {
		fmt.Println(errorText)
	}
	fmt.Println("wifiled v" + VERSION)
	fmt.Println("sends commands to generic wifi led controllers on the network")
	fmt.Println("set environment vars wifiled_ip and wifiled_port")
	fmt.Println("wifiled on -- send on command")
	fmt.Println("wifiled off -- send off command")
	fmt.Println("wifiled dim <RED> <GREEN> <BLUE> -- set RGB values")

}
