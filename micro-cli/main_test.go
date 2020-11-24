package microcli

import (
	"reflect"
	"testing"

	"github.com/go-iot-platform/go-micro"
	"github.com/go-iot-platform/lib/rbac/model"
)

func TestMicro_CallHTTP(t *testing.T) {
	type fields struct {
		Client micro.Client
	}
	type args struct {
		namespace string
		microName string
		endpoint  string
		body      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.ObjectData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Micro{
				Client: tt.fields.Client,
			}
			got, err := m.CallHTTP(tt.args.namespace, tt.args.microName, tt.args.endpoint, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Micro.CallHTTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Micro.CallHTTP() = %v, want %v", got, tt.want)
			}
		})
	}
}
