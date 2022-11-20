package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"wifiled/lib/genericWifiLed"
	"wifiled/lib/toolbox"
)

const VERSION = "0.3"
const BUILD_DATE = "2022-Nov-20"
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
			fmt.Printf("Missing IP - add environment with export %s=x.x.x.x or use -ip=x.x.x.x\n", ENV_IP_KEY)
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
	fmt.Printf("LED: %s:%s\n", ip, port)

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
			// wifiled dim 10
			v := StringToInt(commandLineArguments[2])
			controller.DimTo(v, v, v, v, v)
		} else if commandLineArgumentsLength == 5 {
			// wifiled dim 255 0 200
			redValue, greenValue, blueValue := getRgbFrom(commandLineArguments)
			controller.DimTo(redValue, greenValue, blueValue, 0, 0)
		} else if commandLineArgumentsLength >= 7 {
			// wifiled dim 10 20 30 40 50
			redValue, greenValue, blueValue, warmWhiteValue, coolWhiteValue := getRgbwwFrom(commandLineArguments)
			controller.DimTo(redValue, greenValue, blueValue, warmWhiteValue, coolWhiteValue)
		} else {
			displayHelpText("invalid dim parameters")
		}
	} else if command == "randomize" {
		// wifiled randomize
		// wifiled randomize 0-255
		// wifiled randomize 0-255 255 10-50
		newRed, newGreen, newBlue := 0, 0, 0
		if commandLineArgumentsLength == 3 {
			newRed, newGreen, newBlue = randomizeRGBSingle(commandLineArguments[2])
		} else if commandLineArgumentsLength == 5 {
			newRed, newGreen, newBlue = randomizeRGB(commandLineArguments[2], commandLineArguments[3], commandLineArguments[4])
		} else {
			newRed, newGreen, newBlue = randomizeRGBSingle("0-255")
		}

		fmt.Printf(" Picking Random Color - R%d G%d B%d (#%02X%02X%02X)\n", newRed, newGreen, newBlue, newRed, newGreen, newBlue)
		controller.DimTo(newRed, newGreen, newBlue, 0, 0)

	} else {
		displayHelpText("unknown command")
	}
}

func randomizeRGBSingle(input string) (int, int, int) {
	return randomizeRGB(input, input, input)
}

func randomizeRGB(red string, green string, blue string) (newRed int, newBlue int, newGreen int) {
	rand.Seed(time.Now().Unix())
	offset, lowerBound, foundRange := toolbox.ParseRangeFromString(red, RGBW_MIN, RGBW_MAX)
	if offset > 0 && foundRange {
		newRed = rand.Intn(offset) + lowerBound
	} else {
		newRed = offset
	}
	offset, lowerBound, foundRange = toolbox.ParseRangeFromString(green, RGBW_MIN, RGBW_MAX)
	if offset > 0 && foundRange {
		newBlue = rand.Intn(offset) + lowerBound
	} else {
		newBlue = offset
	}
	offset, lowerBound, foundRange = toolbox.ParseRangeFromString(blue, RGBW_MIN, RGBW_MAX)
	if offset > 0 && foundRange {
		newGreen = rand.Intn(offset) + lowerBound
	} else {
		newGreen = offset
	}
	return
}

func getRgbFrom(commandLineArguments []string) (redValue int, greenValue int, blueValue int) {
	redValue = StringToInt(commandLineArguments[2])
	greenValue = StringToInt(commandLineArguments[3])
	blueValue = StringToInt(commandLineArguments[4])
	return
}
func getRgbwwFrom(commandLineArguments []string) (redValue int, greenValue int, blueValue int, coolWhiteValue int, warmWhiteValue int) {
	redValue = StringToInt(commandLineArguments[2])
	greenValue = StringToInt(commandLineArguments[3])
	blueValue = StringToInt(commandLineArguments[4])
	coolWhiteValue = StringToInt(commandLineArguments[5])
	warmWhiteValue = StringToInt(commandLineArguments[6])
	return
}

func StringToInt(input string) int {
	return toolbox.ConvertStringToBoundedInt(input, RGBW_MAX, RGBW_MIN, RGBW_DEFAULT)
}

func displayHelpText(errorText string) {
	if errorText != "" {
		fmt.Println(errorText)
	}
	fmt.Printf("wifiled v%s (%s)", VERSION, BUILD_DATE)
	fmt.Println(" sends commands to generic wifi led controllers on the network")
	fmt.Println("  wifiled on -- send on command")
	fmt.Println("  wifiled off -- send off command")
	fmt.Println("  wifiled dim <BRIGHTNESS> -- set all channels to value out of 255")
	fmt.Println("  wifiled dim <RED> <GREEN> <BLUE> -- set RGB values out of 255")
	fmt.Println("  wifiled dim <RED> <GREEN> <BLUE> <WARMWHITE> <COOLWHITE> -- set RGBW values out of 255")
	fmt.Println("  wifiled randomize -- sets a random color")
	fmt.Println("  wifiled randomize 0-10 -- sets a random color with values between 0 and 10")
	fmt.Println("  wifiled randomize 255 0-50 0-50 -- sets red fixed and green/blue between 0 and 50")
}
