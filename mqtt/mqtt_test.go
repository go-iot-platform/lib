package mqtt

import (
	// "strconv"
	"testing"
)

func TestContainer(t *testing.T) {
	prefix := "example"
	topic := "example/value"
	result, err := ParseTopic(prefix, topic)
	if err != nil || result == "" {
		t.Fatal("parse topic error")
	}
}
