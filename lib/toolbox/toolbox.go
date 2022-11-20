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

func ParseRangeFromString(commandLineArgument string, min int, max int) (int, int) {
	rgbRange := strings.Split(commandLineArgument, "-")
	offset := 255
	lowerBound := 0
	if len(rgbRange) == 2 {
		upperBound := ConvertStringToBoundedInt(rgbRange[1], max, min, max)
		lowerBound = ConvertStringToBoundedInt(rgbRange[0], max, min, min)
		offset = upperBound - lowerBound
	} else {
		offset = ConvertStringToBoundedInt(commandLineArgument, max, min, max)
	}
	return offset, lowerBound
}
