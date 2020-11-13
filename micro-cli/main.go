package microcli

import (
	"context"
	"encoding/json"
	"fmt"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"github.com/micro/go-micro/client"
	"github.com/go-iot-platform/lib/rbac/model"
	"strings"
)

// Call call micro from endpoint
func Call(namespace string, microName string, endpoint string, body string) (*model.ObjectData, error) {
	var result interface{}
	data := model.ObjectData{}
	err := CallMicro(namespace, microName, endpoint, body, &result)
	if err != nil || result == nil {
		return nil, err
	}
	data.Value = result
	return &data, nil
}

// CallMicro cal micro no protos
func CallMicro(ns string, svcName string, endPoint string, body string, result interface{}) error {
	service := k8s.NewService()
	service.Init()
	c := service.Client()
	t := StrToMap(body)
	request := c.NewRequest(fmt.Sprintf("%s.%s", svcName, ns), endPoint, t, client.WithContentType("application/json"))
	var response json.RawMessage

	if err := c.Call(context.Background(), request, &response); err != nil {
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
func CallWithContext(ctx context.Context, ns string, svcName string, endPoint string, body string, result interface{}) error {
	service := k8s.NewService()
	service.Init()
	c := service.Client()
	t := StrToMap(body)
	request := c.NewRequest(fmt.Sprintf("%s.%s", svcName, ns), endPoint, t, client.WithContentType("application/json"))
	var response json.RawMessage

	if err := c.Call(ctx, request, &response); err != nil {
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
