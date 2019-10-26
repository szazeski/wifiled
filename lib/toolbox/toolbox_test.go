package toolbox

import "testing"

func Test_toolbox_ConvertStringToBoundedInt_blank(t *testing.T) {
	actual := ConvertStringToBoundedInt("",0,0, 0)
	if actual != 0 {
		t.Error("expected to get 0 but got", actual)
	}
}
func Test_toolbox_ConvertStringToBoundedInt_default(t *testing.T) {
	expected := 777
	actual := ConvertStringToBoundedInt("",0,0, expected)
	if actual != expected {
		t.Error("expected to get", expected, "but got", actual)
	}
}
func Test_toolbox_ConvertStringToBoundedInt_floor(t *testing.T) {
	expected := 10
	actual := ConvertStringToBoundedInt("-50",100, expected, 0)
	if actual != expected {
		t.Error("expected to get", expected, "but got", actual)
	}
}