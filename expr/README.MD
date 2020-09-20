# Expr Helper

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://github.com/go-iot-platform/expr)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/go-iot-platform/expr)

ParseText parse template.

## Getting Started!
> Define struct
  ```
  type Data struct {
	Name  string
	Value string
}
```
> Define func want to trigger
```
type TestFn struct {
}

func (self *TestFn) Call(body string) string {
	fmt.Println("mock call function")
	return body
}
```
> How to use
```
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
		// handle error here
	}
```

### Installation

Request handler requires [Go](https://golang.org/) v1.11+ to run.

Install the package.

```sh
$ go get github.com/go-iot-platform/expr
```

#### Kubernetes + Google Cloud

See [KUBERNETES.md](https://github.com/joemccann/dillinger/blob/master/KUBERNETES.md)


### Todos

 - Write MORE Tests
 - Add Night Mode

License
----

MIT