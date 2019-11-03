package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"wifiled/lib/genericWifiLed"
	"wifiled/lib/toolbox"
)

const VERSION = "0.2"
const ENV_IP_KEY = "wifiled_ip"
const ENV_PORT_KEY = "wifiled_port"
const RGBW_MAX = 255
const RGBW_MIN = 0
const RGBW_DEFAULT = RGBW_MIN

func main() {
	commandLineArguments := os.Args
	commandLineArgumentsLength := 0
	for _, value := range commandLineArguments {
		if strings.HasPrefix(value, "-") {
			break
		}
		commandLineArgumentsLength++
	}
	commandLineArguments = commandLineArguments[0:commandLineArgumentsLength]
	command := ""
	if commandLineArgumentsLength >= 2 {
		command = commandLineArguments[1]
	}

	fmt.Println(commandLineArguments)

	commandFlagIp := ""
	flag.StringVar(&commandFlagIp, "ip", "", "set the ip address of the LED Controller")
	commandFlagPort := ""
	flag.StringVar(&commandFlagPort, "port", "", "set the port of the LED Controller")
	flag.Parse()

	if command == "" {
		displayHelpText("")
		return
	}

	ip := os.Getenv(ENV_IP_KEY)
	if ip == "" {
		if commandFlagIp != "" {
			ip = commandFlagIp
		} else {
			displayHelpText("Missing ip environment variable, check your env")
			return
		}
	}
	port := os.Getenv(ENV_PORT_KEY)
	if port == "" {
		if commandFlagPort != "" {
			port = commandFlagPort
		} else {
			port = "5577"
		}
	}
	fmt.Printf("LED: %s:%s", ip, port)

	controller := genericWifiLed.NewController(ip, port)
	if command == "on" {
		controller.On()

	} else if command == "off" {
		controller.Off()

	} else if command == "dim" {
		// wifiled dim R   G   B   WW  CW
		// wifiled dim 255 255 255 255 255
		// 0       1   2   3   4   5   6
		if commandLineArgumentsLength == 3 {
			v := toolbox.ConvertStringToBoundedInt(commandLineArguments[2], RGBW_MAX, RGBW_MIN, RGBW_DEFAULT)
			controller.DimTo(v, v, v, v, v)
		} else if commandLineArgumentsLength == 5 {
			redValue, greenValue, blueValue := getRgbFromCommandLine(commandLineArguments, RGBW_MAX, RGBW_MIN, RGBW_DEFAULT)
			controller.DimTo(redValue, greenValue, blueValue, 0, 0)
		} else if commandLineArgumentsLength >= 7 {
			redValue, greenValue, blueValue := getRgbFromCommandLine(commandLineArguments, RGBW_MAX, RGBW_MIN, RGBW_DEFAULT)
			warmWhiteValue := toolbox.ConvertStringToBoundedInt(commandLineArguments[5], RGBW_MAX, RGBW_MIN, RGBW_DEFAULT)
			coolWhiteValue := toolbox.ConvertStringToBoundedInt(commandLineArguments[6], RGBW_MAX, RGBW_MIN, RGBW_DEFAULT)
			controller.DimTo(redValue, greenValue, blueValue, warmWhiteValue, coolWhiteValue)
		} else {
			displayHelpText("invalid dim parameters")
		}

	} else {
		displayHelpText("unknown command")
	}
}

func getRgbFromCommandLine(commandLineArguments []string, MAX_VALUE int, MIN_VALUE int, DEFAULT_VALUE int) (redValue int, greenValue int, blueValue int) {
	redValue = toolbox.ConvertStringToBoundedInt(commandLineArguments[2], MAX_VALUE, MIN_VALUE, DEFAULT_VALUE)
	greenValue = toolbox.ConvertStringToBoundedInt(commandLineArguments[3], MAX_VALUE, MIN_VALUE, DEFAULT_VALUE)
	blueValue = toolbox.ConvertStringToBoundedInt(commandLineArguments[4], MAX_VALUE, MIN_VALUE, DEFAULT_VALUE)
	return
}

func displayHelpText(errorText string) {
	if errorText != "" {
		fmt.Println(errorText)
	}
	fmt.Println("wifiled v" + VERSION)
	fmt.Println(" sends commands to generic wifi led controllers on the network")
	fmt.Println(" set environment vars wifiled_ip and wifiled_port")
	fmt.Println("  wifiled on -- send on command")
	fmt.Println("  wifiled off -- send off command")
	fmt.Println("  wifiled dim <BRIGHTNESS> -- set all channels to value out of 255")
	fmt.Println("  wifiled dim <RED> <GREEN> <BLUE> -- set RGB values out of 255")
	fmt.Println("  wifiled dim <RED> <GREEN> <BLUE> <WARMWHITE> <COOLWHITE> -- set RGBW values out of 255")
}
