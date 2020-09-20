package redis_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoRedis(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoRedis Suite")
}
