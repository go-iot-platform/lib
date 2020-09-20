package redis_test

import (
	"fmt"
	"github.com/eneoti/dockertest"
	"github.com/go-iot-platform/go-redis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var container *dockertest.Container
var client *redis.Redis
var addr string
var _ = BeforeSuite(func() {
	container, _ = dockertest.RunContainer("redis:5.0-rc-alpine", "redis-test", "6379", func(result string) error {
		addr = result
		return nil
	})
	time.Sleep(time.Millisecond * 1500)
})

var _ = AfterSuite(func() {
	container.Shutdown("redis-test")
	client.Close()
})

var _ = Describe("GoRedis", func() {
	It("Init redis client should receive onconnect event", func() {
		fmt.Printf("addr:%+v\n", addr)
		options := struct {
			Addr *string
		}{
			Addr: &addr,
		}
		client = redis.NewRedis(&options)
		Expect(client.Ping().Err()).NotTo(HaveOccurred())
	})
	It("Set should have no error ", func() {
		_, err := client.Set("test_key", "hello", 0)
		Expect(err).NotTo(HaveOccurred())
	})
	It("Set should receive same with set", func() {
		result, err := client.Get("test_key")
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal("hello"))
	})
	It("Del should have no err", func() {
		err := client.Del("test_key")
		Expect(err).NotTo(HaveOccurred())
	})
	It("SetObject should have no error ", func() {
		type Temp struct {
			Name  string
			Value string
		}
		value := &Temp{Name: "thing1", Value: "hello"}
		_, err := client.SetObject("things", "hello", value)
		Expect(err).NotTo(HaveOccurred())
	})
	It("GetObject should have no error ", func() {
		result := make(map[string]interface{})
		err := client.GetObject("things", "hello", &result)
		Expect(err).NotTo(HaveOccurred())
		Expect(result["Name"]).To(Equal("thing1"))
		err = client.Del("things")
		Expect(err).NotTo(HaveOccurred())
	})
})
