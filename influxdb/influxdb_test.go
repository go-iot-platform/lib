package influxdb_test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"

	"net/http/httptest"

	influxClient "github.com/influxdata/influxdb1-client"
	"github.com/go-iot-platform/influxdb"
)

func TestContainer(t *testing.T) {
	ts := testServer()
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	config := &influxdb.Config{URL: *u}
	query := &influxdb.Query{}
	bps := &influxdb.BatchPoints{}

	influx := testNewClient(t, config)
	if influx != nil {

		testClient_Ping(t, influx)
		testClient_Query(t, influx, query)
		testClient_Write(t, influx, bps)
	}

}
func testNewClient(t *testing.T, config *influxdb.Config) *influxdb.Influx {
	con, err := influxdb.NewClient(config)
	if err != nil {
		t.Fatalf("unexpected error.  expected %v, actual %v", nil, err)
		return nil
	}
	if con == nil {
		t.Fatalf("unexpected error.  expected %s, actual %v", "client.Client", nil)
		return nil
	}
	return con
}
func testClient_Ping(t *testing.T, client *influxdb.Influx) {
	d, version, err := client.Ping()
	if err != nil {
		t.Fatalf("unexpected error.  expected %v, actual %v", nil, err)
	}
	if d.Nanoseconds() == 0 {
		t.Fatalf("expected a duration greater than zero.  actual %v", d.Nanoseconds())
	}
	if version != "x.x" {
		t.Fatalf("unexpected version.  expected %s,  actual %v", "x.x", version)
	}
}
func testClient_Query(t *testing.T, client *influxdb.Influx, query *influxdb.Query) {
	result, err := client.Query(query)
	if err != nil {
		t.Fatalf("unexpected error.  expected %v, actual %v", result, err)
	}
	if result == nil {
		t.Fatalf("unexpected result.  expected %s,  actual %v", "result", nil)
	}
}
func testClient_Write(t *testing.T, client *influxdb.Influx, bps *influxdb.BatchPoints) {
	result, err := client.Write(bps)
	if err != nil {
		t.Fatalf("unexpected error.  expected %v, actual %v", nil, err)
	}
	if result == nil {
		t.Fatalf("unexpected result.  expected %s,  actual %v", "result", nil)
	}
}

// helper functions

func testServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data influxClient.Response
		time.Sleep(50 * time.Millisecond)

		w.Header().Set("X-Influxdb-Version", "x.x")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(data)
	}))
}
