package structs

import (
	// "reflect"
	"fmt"
	"testing"
	"time"

	"github.com/AdrianLungu/decimal"
)

type testStruct struct {
	DataString   string
	DataBool     bool
	DataInteger  int
	DataFloat    float64
	DataTime     time.Time
	DataLargeNum testSubStruct
	DataDecimal  decimal.Decimal
}
type testStructCamelResult struct {
	dataString   string
	dataBool     bool
	dataInteger  int
	dataFloat    float64
	dataTime     time.Time
	dataLargeNum testSubStruct
	dataDecimal  decimal.Decimal
}
type otherStruct struct {
	DataString string
}

type testSubStruct struct {
	IntPart int
	IsZero  bool
}

func TestContainer(t *testing.T) {

	to := testStruct{
		DataString:  "default string",
		DataBool:    true,
		DataInteger: 100,
		DataFloat:   float64(1.001112),
		DataTime:    time.Now(),
		DataLargeNum: testSubStruct{
			IntPart: 1000,
			IsZero:  false,
		},
		DataDecimal: decimal.NewFromFloat(1000),
	}
	from := otherStruct{
		DataString: "another string",
	}
	// testMergeOverwrite(t, from, to)
	// testMergeOverwriteCamel(t, from, to)
	testMerge(t, &from)
	testMergeToMap(t, to, from)
}

func testMergeOverwrite(t *testing.T, from otherStruct, to testStruct) {
	var target testStruct
	if err := MergeOverwrite(to, from, &target); err != nil {
		t.Fatal(err)
	}
	if expected := from.DataString; target.DataString != expected {
		t.Errorf("want %s got %s", expected, target.DataString)
	}
}
func testMergeOverwriteCamel(t *testing.T, from otherStruct, to testStruct) {
	var target testStructCamelResult
	if err := MergeOverwriteCamel(to, from, &target); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("target : %+v\n", target)
	if expected := from.DataString; target.dataString != expected {
		t.Errorf("want %s got %s", expected, target.dataString)
	}
}
func testMerge(t *testing.T, from *otherStruct) {
	target := new(testStruct)
	Merge(target, from)
	if expected := from.DataString; target.DataString != expected {
		t.Errorf("want %s got %s", expected, target.DataString)
	}
}
func testMergeToMap(t *testing.T, to testStruct, from interface{}) {
	_, err := MergeToMap(to, from)
	if err != nil {
		t.Fatal(err)
	}
}
func testConvertStringToMap(t *testing.T, str string) {
	result := ConvertStringToMap(str)
	fmt.Printf("convert to string map%v\n", result)
}
