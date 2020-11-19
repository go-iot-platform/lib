# Go-redis Helper

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://github.com/go-iot-platform/lib/go-redis)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/go-iot-platform/lib/go-redis)

Redis client for golang

## Getting Started!

> Get by key
  ```
  func (r *Redis) Get(key string) (string, error) {
    return r.db.Get(key).Result()
  }
  ```
> Get by key and field
  ```
  func (r *Redis) GetObject(objectKey string, field string, result interface{}) error {
    temp, err := r.db.HGet(objectKey, field).Bytes()
    if err != nil {
      return err
    }
    err = msgpack.Unmarshal(temp, result)
    return err
  }
  ```
> Set by key
  ```
  func (r *Redis) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return r.db.Set(key, value, expiration).Result()
  }
  ```
> Set by key and field
  ```
  func (r *Redis) SetObject(objectKey string, field string, value interface{}) (bool, error) {
	bytes, err := msgpack.Marshal(value)
	if err != nil {
		return false, err
	}
	return r.db.HSet(objectKey, field, bytes).Result()
  }
  ```
> Check exist by key and field
  ```
  // CheckExistedObject will return true if the object is existed.
  func (r *Redis) CheckExistedObject(objectKey string, field string) (bool, error) {
    existed, err := r.db.HExists(objectKey, field).Result()
    if err != nil {
      return false, err
    }
    if existed == true {
      return true, nil
    }
    return false, nil
  }
  ```
### Installation

Request handler requires [Go](https://golang.org/) v1.13+ to run.

Install the package.

```sh
$ go get -u github.com/go-iot-platform/lib/go-redis
```

#### Kubernetes + Google Cloud

See [KUBERNETES.md](https://github.com/joemccann/dillinger/blob/master/KUBERNETES.md)


### Todos

 - Write MORE Tests
 - Add Night Mode

License
----

MIT