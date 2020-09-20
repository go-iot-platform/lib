package expr

import (
	"fmt"

	// "strconv"
	"testing"
)

type TestFn struct {
}
type Data struct {
	Name  string
	Value string
}

// Call test func
func (self *TestFn) Call(body string) string {
	fmt.Println("mock call function")
	return body
}

func TestContainer(t *testing.T) {
	test := new(TestFn)
	data := Data{
		Name:  "Nguyen Cat Pham",
		Value: "oscar",
	}
	obj := map[string]interface{}{
		"data": data,
		"fn":   test,
	}
	strParsing := `{{$x := printf "{\"serial\":%q}" "response data" }}{{.fn.Call $x}}`
	result, err := ParseText("test", strParsing, obj)
	if err != nil || result == "" {
		t.Fatal("parser error")
	}
}
