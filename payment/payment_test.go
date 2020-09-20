package payment

import (
	"testing"
	"time"

	"github.com/go-iot-platform/go-redis"
	"github.com/go-iot-platform/rbac/model"
)

// SubscriptionKey Seperate Char
const SubscriptionKey = "test"

func TestContainer(t *testing.T) {
	cn := "123456"
	sn := "test"
	now := time.Now()
	options := Redis{
		Addr: "192.168.0.53:6379",
	}
	// model := &payment.SubscriptionCache{
	// 	UUID:             "123456",
	// 	PackageID:        "123456",
	// 	StartDate:        now,
	// 	EndDate:          now,
	// 	Type:             2,
	// 	Meta:             nil,
	// 	Duration:         0,
	// 	IntervalTime:     0,
	// 	Quota:            60,
	// 	OldPrice:        53500000,
	// 	CustomerNumber:   "123456",
	// 	Service:          "test",
	// 	Status:           2,
	// 	HaveTrialPackage: true,
	// 	TrialDuration:    10,
	// }
	model := &model.SubscriptionCache{
		UUID:             "123456",
		PackageID:        "123456",
		StartDate:        now,
		EndDate:          now.Add(time.Hour * 10 * 24 * 30),
		Type:             3,
		Meta:             nil,
		Duration:         12 * 30,
		IntervalTime:     2,
		Quota:            60,
		OldPrice:         53500000,
		CustomerNumber:   "123456",
		Service:          "test",
		Status:           2,
		HaveTrialPackage: false,
		TrialDuration:    "1-m",
	}
	client := redis.NewRedis(&options)

	client.SetObject("onsky:payment:subscriptions", cn+sn, model)

	p := Init(client, SubscriptionKey)
	testValidatePayment(p, cn, sn, t)
}
func testValidatePayment(p *Payment, cn string, sn string, t *testing.T) {
	duration, err := p.Validate(cn, sn)
	if err != nil {
		t.Fatalf("payment reject: %+v\n", err)
	}
	t.Logf("duration:%+v\n", duration)
}

// Redis configuration
type Redis struct {
	Addr     string `yaml:"addr,omitempty"`
	Password string `yaml:"password,omitempty"`
}
