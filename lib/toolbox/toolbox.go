package toolbox

import "strconv"

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