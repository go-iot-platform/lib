package microcli

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-iot-platform/lib/rbac/model"
)

func TestCall(t *testing.T) {
	type args struct {
		namespace string
		microName string
		endpoint  string
		body      string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.ObjectData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Call(tt.args.namespace, tt.args.microName, tt.args.endpoint, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Call() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCallMicro(t *testing.T) {
	type args struct {
		ns       string
		svcName  string
		endPoint string
		body     string
		result   interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CallMicro(tt.args.ns, tt.args.svcName, tt.args.endPoint, tt.args.body, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("CallMicro() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCallWithContext(t *testing.T) {
	type args struct {
		ctx      context.Context
		ns       string
		svcName  string
		endPoint string
		body     string
		result   interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CallWithContext(tt.args.ctx, tt.args.ns, tt.args.svcName, tt.args.endPoint, tt.args.body, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("CallWithContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStrToMap(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToMap(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
