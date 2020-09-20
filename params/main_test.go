package params

import (
	"testing"
)

func TestContainer(t *testing.T) {
	paramInt := "1,2,3,4,5"
	paramString := "one,two,three"
	sep := ","

	testConvertToInt(paramInt, sep, t)
	testConvertToString(paramString, sep, t)
}
func testConvertToInt(param string, sep string, t *testing.T) {
	result, err := ConvertToInt(param, sep)
	if err != nil || result == nil {
		t.Fatal("convert error")
	}
}
func testConvertToString(param string, sep string, t *testing.T) {
	result, err := ConvertToString(param, sep)
	if err != nil || result == nil {
		t.Fatal("convert error")
	}
}
