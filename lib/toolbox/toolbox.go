package toolbox

import (
	"strconv"
	"strings"
)

func ConvertStringToBoundedInt(input string, max int, min int, defaultValue int) int {
	value, err := strconv.Atoi(input)
	if err != nil {
		return defaultValue
	}
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}

func ParseRangeFromString(commandLineArgument string, min int, max int) (offset int, lowerBound int, foundRange bool) {
	rgbRange := strings.Split(commandLineArgument, "-")
	if len(rgbRange) == 2 {
		upperBound := ConvertStringToBoundedInt(rgbRange[1], max, min, max)
		lowerBound = ConvertStringToBoundedInt(rgbRange[0], max, min, min)
		offset = upperBound - lowerBound
		foundRange = true
		return
	} else {
		offset = ConvertStringToBoundedInt(commandLineArgument, max, min, max)
		return
	}
}

func ParseHexColor(s string) (int, int, int, error) {
	s = strings.TrimPrefix(s, "#")
	red, err := strconv.ParseInt(s[0:2], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}
	green, err := strconv.ParseInt(s[2:4], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}
	blue, err := strconv.ParseInt(s[4:6], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}
	return int(red), int(green), int(blue), nil
}
