# Names Helper

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://github.com/go-iot-platform/names)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/go-iot-platform/names)

Names parser helper.

## Getting Started!

> get first name
```
result := GetFirstName(name)
	if result == "" {
		// handle error here
	}
	return result
```

> get last name
```
result := GetLastName(name)
	if result == "" {
		// handle error here
	}
	return result
```


> get full name
```
result := GetFullName(firstname,lastname)
	if result == "" {
		// handle error here
	}
	return result
```


### Installation

Request handler requires [Go](https://golang.org/) v1.11+ to run.

Install the package.

```sh
$ go get github.com/go-iot-platform/names
```

#### Kubernetes + Google Cloud

See [KUBERNETES.md](https://github.com/joemccann/dillinger/blob/master/KUBERNETES.md)


### Todos

 - Write MORE Tests
 - Add Night Mode

License
----

MIT