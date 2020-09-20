# I/O Helper

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://github.com/go-iot-platform/io)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/go-iot-platform/io)

Package io implements some I/O utility functions.


# Getting Started!

> Read all file in a directory
```
func testReadAll(path string, t *testing.T) string {
	result, err := ReadAll(path)
	if err != nil || result == "" {
		t.Fatal("read file error")
	}
	return result
}
```
> Write to a file already exist
```
func testWriteFile(path string, json string, t *testing.T) {
	err := WriteFile(path, json)
	if err != nil {
		t.Log(err)
		t.Fatal("write file error")
	}
}
```
> Read from directory
```
func testReadDir(path string, t *testing.T) {
	_, err := ReadDir(path)
	if err != nil {
		t.Fatal("read dirctory error")
	}
}
```

### Installation

Request handler requires [Go](https://golang.org/) v1.11+ to run.

Install the package.

```sh
$ go get github.com/go-iot-platform/io
```

#### Kubernetes + Google Cloud

See [KUBERNETES.md](https://github.com/joemccann/dillinger/blob/master/KUBERNETES.md)


### Todos

 - Write MORE Tests
 - Add Night Mode

License
----

MIT
