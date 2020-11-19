package payment

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/go-iot-platform/lib/errors"
	"github.com/go-iot-platform/lib/go-redis"
	"github.com/go-iot-platform/lib/rbac/model"
)

// ServiceName Seperate Char
const ServiceName = "payment"

const (
	Time     = iota + 1 // quota + time => duration>0, duration = endate - startdate
	Quota               // quota only => duration =0 => depend on service, nerver expired
	Interval            // quota + time , duration = (endate - startdate)/interval time
)

// Payment struct receiver
type Payment struct {
	DB              *redis.Redis
	SubscriptionKey string
}

// Init func
func Init(db *redis.Redis, subKey string) *Payment {
	return &Payment{
		DB:              db,
		SubscriptionKey: subKey,
	}
}

// Validate payment subscription
func (p *Payment) Validate(customerNumber string, serviceName string) (uint16, error) {
	now := time.Now()
	if p == nil || customerNumber == "" || serviceName == "" {
		return 0, errors.BadRequest(ServiceName, "validate:invalidParams")
	}
	subscription, err := getSubscription(p, customerNumber, serviceName)
	if err != nil {
		return 0, err
	}
	duration, err := checkSubscription(subscription, now)
	if err != nil {
		return 0, err
	}
	if duration == 0 {
		return 0, errors.Forbidden(ServiceName, "validate:expired")
	}
	return duration, nil
}
func checkSubscription(subscription *model.SubscriptionCache, now time.Time) (uint16, error) {
	if subscription.Status != 2 { //hardcode
		return 0, errors.Forbidden(ServiceName, "validate:inactiveSubscription")
	}
	if subscription.OldPrice <= 0 {
		return 0, errors.Forbidden(ServiceName, "validate:invalidPrice")
	}
	if subscription.Quota <= 0 {
		return 0, errors.Forbidden(ServiceName, "validate:invalidQuota")
	}
	switch subscription.Type {
	case Quota:
		if subscription.Duration != 0 {
			return 0, errors.Forbidden(ServiceName, "validate:invalidDuration")
		}
		if subscription.IntervalTime != 0 {
			return 0, errors.Forbidden(ServiceName, "validate:invalidIntervalTime")
		}
		quota := subscription.Quota
		if subscription.HaveTrialPackage {
			trialDuration, err := strconv.Atoi(subscription.TrialDuration)
			if err != nil {
				return 0, errors.Forbidden(ServiceName, "validate:invalidTrialDuration")
			}
			if trialDuration <= 0 {
				return 0, errors.Forbidden(ServiceName, "validate:invalidTrialDuration")
			}
			quota += uint16(trialDuration)
		}
		return quota, nil
	case Time:
		if subscription.StartDate.After(now) || subscription.EndDate.Before(now) || subscription.EndDate.Before(subscription.StartDate) {
			return 0, errors.Forbidden(ServiceName, "validate:invalidDate")
		}
		if subscription.IntervalTime != 0 {
			return 0, errors.Forbidden(ServiceName, "validate:invalidIntervalTime")
		}
		if subscription.Duration <= 0 {
			return 0, errors.Forbidden(ServiceName, "validate:invalidDuration")
		}
		// duration := subscription.Duration
		// if subscription.HaveTrialPackage {
		// 	trialDuration, err := CalculateTrial(subscription.StartDate, subscription.TrialDuration)
		// 	if err != nil {
		// 		return 0, err
		// 	}
		// 	if subscription.TrialDuration == 0 {
		// 		return 0, errors.Forbidden(ServiceName, "validate:invalidTrialDuration")
		// 	}
		// 	duration += subscription.TrialDuration
		// }
		if subscription.Duration > uint16(subscription.EndDate.Sub(subscription.StartDate).Hours()) {
			return 0, errors.Forbidden(ServiceName, "validate:invalidDuration")
		}
		return subscription.Quota, nil
	case Interval:
		if subscription.StartDate.After(now) || subscription.EndDate.Before(now) || subscription.EndDate.Before(subscription.StartDate) {
			return 0, errors.Forbidden(ServiceName, "validate:invalidDate")
		}
		if subscription.IntervalTime <= 0 {
			return 0, errors.Forbidden(ServiceName, "validate:invalidIntervalTime")
		}
		rangeDate := int(subscription.EndDate.Sub(subscription.StartDate).Hours())
		day := now.Day()
		invervalHour := int(subscription.IntervalTime) * day * 24
		duration := rangeDate / invervalHour
		if subscription.Duration <= 0 || int(subscription.Duration) != duration {
			return 0, errors.Forbidden(ServiceName, "validate:invalidDuration")
		}
		return subscription.Quota, nil
	default:
		return 0, errors.Forbidden(ServiceName, "validate:invalidType")
	}
	return 0, errors.Forbidden(ServiceName, "validate:invalidType")
}
func getSubscription(p *Payment, customerNumber string, serviceName string) (*model.SubscriptionCache, error) {
	var subscription = new(model.SubscriptionCache)
	var t interface{}
	// get subscription from cache
	err := p.DB.GetObject(p.SubscriptionKey, customerNumber+serviceName, subscription)
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, errors.Forbidden(ServiceName, "validate:unregestered")
		}
		return nil, err
	}
	buffer, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(buffer, subscription)
	if subscription == nil {
		return nil, errors.InternalServerError(ServiceName, "validate:wrongData")
	}
	if subscription.CustomerNumber != customerNumber || subscription.Service != serviceName {
		return nil, errors.BadRequest(ServiceName, "validate:wrongCustomer")
	}
	return subscription, nil
}

// CalculateTrial cal trial package duration
func CalculateTrial(st time.Time, d string) (time.Time, error) {
	s := strings.Split(d, "-")
	v, err := strconv.Atoi(s[0])
	if err != nil {
		return time.Time{}, err
	}
	switch s[1] {
	case "m":
		return st.AddDate(0, v, 0), nil
	case "w":
		return st.AddDate(0, 0, v*7), nil
	case "d":
		return st.AddDate(0, 0, v), nil
	case "y":
		return st.AddDate(v, 0, 0), nil
	case "mi": //! This case is for testing
		return st.Add(time.Duration(v) * time.Minute), nil
	default:
		return time.Time{}, nil
	}
}
