package toolbox

import (
	"math/rand"
	"strconv"
	"testing"
)

func Test_toolbox_ConvertStringToBoundedInt_blank(t *testing.T) {
	actual := ConvertStringToBoundedInt("", 0, 0, 0)
	if actual != 0 {
		t.Error("expected to get 0 but got", actual)
	}
}
func Test_toolbox_ConvertStringToBoundedInt_default(t *testing.T) {
	expected := 777
	actual := ConvertStringToBoundedInt("", 0, 0, expected)
	if actual != expected {
		t.Error("expected to get", expected, "but got", actual)
	}
}
func Test_toolbox_ConvertStringToBoundedInt_floor(t *testing.T) {
	expected := 10
	actual := ConvertStringToBoundedInt("-50", 100, expected, 0)
	if actual != expected {
		t.Error("expected to get", expected, "but got", actual)
	}
}

func Test_toolbox_ParseRangeFromString(t *testing.T) {
	TryParseRangeFromString(t, "", 0, 255, false)
	TryParseRangeFromString(t, "10-50", 10, 40, true)
	TryParseRangeFromString(t, "123", 0, 123, false)
	TryParseRangeFromString(t, "0", 0, 0, false)
}

func TryParseRangeFromString(t *testing.T, input string, expectedLowerBound int, expectedOffset int, expectedRange bool) {
	actualOffset, actualLowerBound, actualFound := ParseRangeFromString(input, 0, 255)

	if actualOffset != expectedOffset {
		t.Errorf("'%s' expected Offest to get %d but got %d", input, expectedOffset, actualOffset)
	}
	if actualLowerBound != expectedLowerBound {
		t.Errorf("'%s' expected LowerBound to get %d but got %d", input, expectedLowerBound, actualLowerBound)
	}

	if actualFound != expectedRange {
		t.Errorf("'%s' expected range to get %v but got %v", input, expectedRange, actualFound)

	}
}

func Test_toolbox_ParseRangeFromString_to_RandomizeRgb(t *testing.T) {

	for i := 0; i < 100; i++ {
		offset, lowerBound, _ := ParseRangeFromString("20-50", 0, 255)
		actual := rand.Intn(offset) + lowerBound

		if actual < 20 && actual > 50 {
			t.Error("out of expected range 20-50 but got", actual)
		}
	}
}

func rgbToHex(r, g, b int) string {
	return strconv.FormatInt(int64(r), 16) + strconv.FormatInt(int64(g), 16) + strconv.FormatInt(int64(b), 16)
}

func Test_toolbox_ParseHexColor(t *testing.T) {
	expectedRed := 18
	expectedGreen := 52
	expectedBlue := 86
	actualRed, actualGreen, actualBlue, err := ParseHexColor(rgbToHex(expectedRed, expectedGreen, expectedBlue))
	if err != nil {
		t.Error("unexpected error", err)
	}
	if actualRed != expectedRed {
		t.Error("expected to get", expectedRed, "but got", actualRed)
	}
	if actualGreen != expectedGreen {
		t.Error("expected to get", expectedGreen, "but got", actualGreen)
	}
	if actualBlue != expectedBlue {
		t.Error("expected to get", expectedBlue, "but got", actualBlue)
	}
}

func Test_toolbox_ParseHexColor_InvalidInput(t *testing.T) {
	r, g, b, err := ParseHexColor("invalid")
	if err == nil {
		t.Error("expected error but got nil")
	}
	if r != 0 || g != 0 || b != 0 {
		t.Error("expected to get 0 for all rgb values but got", r, g, b)
	}

}
