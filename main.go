package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"wifiled/lib/genericWifiLed"
	"wifiled/lib/toolbox"
)

const VERSION = "0.4"
const BUILD_DATE = "2023-Dec-8"

const KEY_ENV_PREFIX = "wifiled_"
const KEY_IP = "ip"
const KEY_PORT = "port"
const KEY_TIMEOUT = "timeout"
const KEY_AVOID_WHITE = "avoidwhite"

const RGBW_MAX = 255
const RGBW_MIN = 0
const RGBW_DEFAULT = RGBW_MIN

const AVOID_WHITE_THRESHOLD = 150

func main() {
	commandLineArguments := os.Args
	commandLineArgumentsLength := 0
	commandIndex := 0
	commandFlagIp := ""
	commandFlagPort := ""
	timeout := ""
	avoidWhite := false
	for i, value := range commandLineArguments {
		if strings.HasPrefix(value, "-") {
			value = strings.ToLower(value)
			value = strings.Replace(value, "--", "-", 1)
			if strings.HasPrefix(value, "-"+KEY_TIMEOUT) {
				timeout = strings.Replace(value, "-timeout=", "", 1)
			}
			if strings.HasPrefix(value, "-"+KEY_IP) {
				commandFlagIp = strings.Replace(value, "-ip=", "", 1)
			}
			if strings.HasPrefix(value, "-"+KEY_PORT) {
				commandFlagPort = strings.Replace(value, "-port=", "", 1)
			}
			if strings.HasPrefix(value, "-"+KEY_AVOID_WHITE) {
				avoidWhite = true
			}
			continue
		}
		commandLineArgumentsLength++
		commandIndex = i
	}
	command := ""
	if commandIndex > 0 {
		command = commandLineArguments[commandIndex]
	}

	if command == "" {
		displayHelpText("")
		return
	}

	ip := env(KEY_IP)
	if ip == "" {
		if commandFlagIp != "" {
			ip = commandFlagIp
		} else {
			fmt.Printf("Missing IP - add environment with export %s=x.x.x.x or use -ip=\"x.x.x.x\"\n", KEY_ENV_PREFIX+KEY_IP)
			return
		}
	}
	port := env(KEY_PORT)
	if port == "" {
		if commandFlagPort != "" {
			port = commandFlagPort
		}
	}
	if env(KEY_AVOID_WHITE) != "" {
		avoidWhite = true
	}

	timeoutInt := toolbox.ConvertStringToBoundedInt(timeout, 60, 1, 5)
	controller := genericWifiLed.NewController(ip, port, timeoutInt)
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
		if avoidWhite && newRed > AVOID_WHITE_THRESHOLD && newGreen > AVOID_WHITE_THRESHOLD && newBlue > AVOID_WHITE_THRESHOLD {
			fmt.Println(" Avoiding White")
			newBlue = 0
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
	offset, lowerBound, foundRange := toolbox.ParseRangeFromString(red, RGBW_MIN, RGBW_MAX)
	if offset > 0 && foundRange {
		newRed = rand.Intn(offset) + lowerBound
	} else {
		newRed = offset
	}
	offset, lowerBound, foundRange = toolbox.ParseRangeFromString(green, RGBW_MIN, RGBW_MAX)
	if offset > 0 && foundRange {
		newGreen = rand.Intn(offset) + lowerBound
	} else {
		newGreen = offset
	}
	offset, lowerBound, foundRange = toolbox.ParseRangeFromString(blue, RGBW_MIN, RGBW_MAX)
	if offset > 0 && foundRange {
		newBlue = rand.Intn(offset) + lowerBound
	} else {
		newBlue = offset
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
func env(key string) string {
	return os.Getenv(KEY_ENV_PREFIX + key)
}

func displayHelpText(errorText string) {
	if errorText != "" {
		fmt.Println(errorText)
	}
	fmt.Printf("wifiled v%s (%s)", VERSION, BUILD_DATE)
	fmt.Println(" sends commands to generic wifi led controllers on the network")
	fmt.Println("  wifiled -ip=\"1.2.3.4,5.6.7.8\" -port=5577 -timeout=1 -avoidwhite off")
	fmt.Println("  wifiled on -- send on command")
	fmt.Println("  wifiled off -- send off command")
	fmt.Println("  wifiled dim <BRIGHTNESS> -- set all channels to value out of 255")
	fmt.Println("  wifiled dim <RED> <GREEN> <BLUE> -- set RGB values out of 255")
	fmt.Println("  wifiled dim <RED> <GREEN> <BLUE> <WARMWHITE> <COOLWHITE> -- set RGBW values out of 255")
	fmt.Println("  wifiled randomize -avoidwhite -- sets a random color")
	fmt.Println("  wifiled randomize 0-10 -- sets a random color with values between 0 and 10")
	fmt.Println("  wifiled randomize 255 0-50 0-50 -- sets red fixed and green/blue between 0 and 50")
}
