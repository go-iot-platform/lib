# micro-cli Helper

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://github.com/go-iot-platform/micro-cli)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/go-iot-platform/micro-cli)

micro-cli parser helper.

## Getting Started!

> How to use
```
    prefix := "example"
	topic := "example/value"
	result, err := ParseTopic(prefix, topic)
	if err != nil || result == "" {
		t.Fatal("parse topic error")
	}
```

### Installation

Request handler requires [Go](https://golang.org/) v1.11+ to run.

Install the package.

```sh
$ go get github.com/go-iot-platform/micro-cli
```

#### Kubernetes + Google Cloud

See [KUBERNETES.md](https://github.com/joemccann/dillinger/blob/master/KUBERNETES.md)


### Todos

 - Write MORE Tests
 - Add Night Mode

License
----

MIT