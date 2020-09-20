# Version Helper

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://github.com/go-iot-platform)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/go-iot-platform)

Auto versioning update helper.

## Getting Started!

```
oldVersion := "1.1.2.1.9.5"
result, err := UpdateVersion(oldVersion)
	if err != nil || result == "" {
		// handle error here
	}
	return result
```


### Installation

Request handler requires [Go](https://golang.org/) v1.11+ to run.

Install the package.

```sh
$ go get github.com/go-iot-platform/version
```

#### Kubernetes + Google Cloud

See [KUBERNETES.md](https://github.com/joemccann/dillinger/blob/master/KUBERNETES.md)


### Todos

 - Write MORE Tests
 - Add Night Mode

License
----

MIT