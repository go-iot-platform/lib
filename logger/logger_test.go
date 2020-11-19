package logger_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-iot-platform/lib/logger"
)

func TestContainer(t *testing.T) {
	logLevel := 0
	logLocation := ""
	logBucket := ""
	serviceName := ""

	testLog(t, serviceName, logLevel, logLocation, logBucket)

}
func testLog(t *testing.T, serviceName string, logLevel int, logLocation string, logBucket string) {
	logger, file, err := logger.Init(serviceName, logLevel, logLocation, logBucket)
	if err != nil || logger == nil {
		fmt.Printf("err:%+v\n", err)
		t.Fatal("log fail")
	}
	defer file.Close()
	err = errors.New("some error")
	teststr := Tester{
		Name:   "test",
		Number: 10,
	}

	logger.Log().Interface("test", teststr).Err(err).Msg("test warn")
	logger.Debug().Err(err).Interface("test", teststr).Msg("test debug")
	logger.Err(err).Interface("test", teststr).Msg("test error")
}

type Tester struct {
	Name   string
	Number int
}
