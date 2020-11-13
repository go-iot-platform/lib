# Struct Helper

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://github.com/nguyencatpham/request-handler)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/nguyencatpham/request-handler)

Merge two different structs, and overwrite fields

## What is this

To make creating struct for tests easier.

## Getting Started!

  ```
    to := testStruct{
		String:  "default string",
		Bool:    true,
		Integer: 100,
		Float:   float64(1.001112),
		Time:    time.Now(),
		LargeNum: testSubStruct{
			IntPart: 1000,
			IsZero:  false,
		},
		Decimal: decimal.NewFromFloat(1000),
	}
	from := otherStruct{
		String: "another string",
	}
	var target testStruct
	if err := MergeOverwrite(to, from, &target); err != nil {
		t.Fatal(err)
	}
  ```
### Other function

#### Merge
> Merge receives two POINTER of structs, and merges them excluding fields with tag name: `structs`, value "-"
```
Merge(dst, src interface{})
```
#### MergeOverwrite
> Merge 2 structs overwriten
```
MergeOverwrite(to, from, dst interface{})
```
#### MergeOverwriteCamel (not stable)
> Merge 2 structs overwriten supported cammel case
```
MergeOverwriteCamel(to, from, dst interface{})
```
#### MergeToMap
> Merge 2 structs to map
```
MergeToMap(to, from interface{}) (map[string]interface{}, error)
```
#### ConvertStringToMap
> Convert string json to map
```
ConvertStringToMap(in string) map[string]interface{}
```

### Installation

Request handler requires [Go](https://golang.org/) v1.11+ to run.

Install the package.

```sh
$ go get github.com/go-iot-platform/lib/structs
```

#### Kubernetes + Google Cloud

See [KUBERNETES.md](https://github.com/joemccann/dillinger/blob/master/KUBERNETES.md)


### Todos

 - Write MORE Tests
 - Add Night Mode

License
----

MIT
