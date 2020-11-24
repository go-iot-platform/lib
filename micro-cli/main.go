package microcli

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-iot-platform/go-micro"
	"github.com/go-iot-platform/go-micro/client"
	"github.com/go-iot-platform/go-micro/client/selector/static"
	"github.com/go-iot-platform/go-micro/service/grpc"
	"github.com/go-iot-platform/lib/rbac/model"
	k8s "github.com/go-iot-platform/micro-plugins/registry/kubernetes"
)

const serviceName = "CLI"

// Micro represent micro struct
type Micro struct {
	micro.Client
}

// Init the micro service
func Init() Micro {
	options := []micro.Option{
		micro.Name(serviceName),
		micro.RegisterTTL(time.Second * 30),
		micro.RegisterInterval(time.Second * 10),
		micro.Registry(k8s.NewRegistry()),
		micro.Selector(static.NewSelector()),
	}
	service := grpc.NewService(options...)
	service.Init()
	return service.Client()
}

// CallHTTP call micro from endpoint
func (m *Micro) CallHTTP(namespace string, microName string, endpoint string, body string) (*model.ObjectData, error) {
	var result interface{}
	data := model.ObjectData{}
	err := m.CallMicro(namespace, microName, endpoint, body, &result)
	if err != nil || result == nil {
		return nil, err
	}
	data.Value = result
	return &data, nil
}

// CallMicro cal micro no protos
func (m *Micro) CallMicro(ns string, svcName string, endPoint string, body string, result interface{}) error {
	t := StrToMap(body)
	request := m.NewRequest(fmt.Sprintf("%s.%s", svcName, ns), endPoint, t, client.WithContentType("application/json"))
	var response json.RawMessage

	if err := m.Call(context.Background(), request, &response); err != nil {
		return err
	}
	// raw := StrToMap(string(response))
	raw := string(response)
	if raw != "" {
		// err := mapstructure.Decode(raw, &result)
		err := json.Unmarshal([]byte(raw), result)
		if err != nil {
			return err
		}
	}
	return nil
}

// CallWithContext cal micro no protos with context
func (m *Micro) CallWithContext(ctx context.Context, ns string, svcName string, endPoint string, body string, result interface{}) error {
	t := StrToMap(body)
	request := m.NewRequest(fmt.Sprintf("%s.%s", svcName, ns), endPoint, t, client.WithContentType("application/json"))
	var response json.RawMessage

	if err := m.Call(ctx, request, &response); err != nil {
		return err
	}
	// raw := StrToMap(string(response))
	raw := string(response)
	if raw != "" {
		// err := mapstructure.Decode(raw, &result)
		err := json.Unmarshal([]byte(raw), result)
		if err != nil {
			return err
		}
	}
	return nil
}

// StrToMap string to map
func StrToMap(in string) map[string]interface{} {
	var result map[string]interface{}
	d := json.NewDecoder(strings.NewReader(in))
	d.UseNumber()

	if err := d.Decode(&result); err != nil {
		return nil
	}
	return result
}
