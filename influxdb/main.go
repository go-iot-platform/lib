package influxdb

import (
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/go-iot-platform/lib/structs"

	client "github.com/influxdata/influxdb1-client"
)

// Config is used to specify what server to connect to.
// URL: The URL of the server connecting to.
// Username/Password are optional. They will be passed via basic auth if provided.
// UserAgent: If not provided, will default "InfluxDBClient",
// Timeout: If not provided, will default to 0 (no timeout)
type Config struct {
	URL              url.URL
	UnixSocket       string
	Username         string
	Password         string
	UserAgent        string
	Timeout          time.Duration
	Precision        string
	WriteConsistency string
	UnsafeSsl        bool
	Proxy            func(req *http.Request) (*url.URL, error)
	TLS              *tls.Config
}

// Query is used to send a command to the server. Both Command and Database are required.
type Query struct {
	Command  string
	Database string

	// RetentionPolicy tells the server which retention policy to use by default.
	// This option is only effective when querying a server of version 1.6.0 or later.
	RetentionPolicy string

	// Chunked tells the server to send back chunked responses. This places
	// less load on the server by sending back chunks of the response rather
	// than waiting for the entire response all at once.
	Chunked bool

	// ChunkSize sets the maximum number of rows that will be returned per
	// chunk. Chunks are either divided based on their series or if they hit
	// the chunk size limit.
	//
	// Chunked must be set to true for this option to be used.
	ChunkSize int

	// NodeID sets the data node to use for the query results. This option only
	// has any effect in the enterprise version of the software where there can be
	// more than one data node and is primarily useful for analyzing differences in
	// data. The default behavior is to automatically select the appropriate data
	// nodes to retrieve all of the data. On a database where the number of data nodes
	// is greater than the replication factor, it is expected that setting this option
	// will only retrieve partial data.
	NodeID int
}

// BatchPoints is used to send batched data in a single write.
// Database and Points are required
// If no retention policy is specified, it will use the databases default retention policy.
// If tags are specified, they will be "merged" with all points. If a point already has that tag, it will be ignored.
// If time is specified, it will be applied to any point with an empty time.
// Precision can be specified if the time is in epoch format (integer).
// Valid values for Precision are n, u, ms, s, m, and h
type BatchPoints struct {
	Points           []client.Point    `json:"points,omitempty"`
	Database         string            `json:"database,omitempty"`
	RetentionPolicy  string            `json:"retentionPolicy,omitempty"`
	Tags             map[string]string `json:"tags,omitempty"`
	Time             time.Time         `json:"time,omitempty"`
	Precision        string            `json:"precision,omitempty"`
	WriteConsistency string            `json:"-"`
}

// Influx model
type Influx struct {
	db *client.Client
}

// DB constructor
func (influx *Influx) DB() *client.Client {
	return influx.db
}

// NewClient constructor
func NewClient(config *Config) (*Influx, error) {
	conf := new(client.Config)
	structs.Merge(conf, config)
	con, err := client.NewClient(*conf)
	if err != nil {
		return nil, err
	}
	return &Influx{
		db: con,
	}, nil
}

// Ping connectivity
func (influx *Influx) Ping() (time.Duration, string, error) {
	return influx.db.Ping()
}

// Query query
func (influx *Influx) Query(query *Query) (*[]client.Result, error) {
	q := new(client.Query)
	structs.Merge(q, query)
	response, err := influx.db.Query(*q)
	if err != nil {
		return nil, err
	}
	if response.Error() != nil {
		return nil, response.Error()
	}
	return &response.Results, nil
}
func (influx *Influx) Write(batchPoints *BatchPoints) (*[]client.Result, error) {
	if batchPoints == nil {
		return nil, errors.New("influxdb:write:emptyBatchPoint")
	}
	bp := new(client.BatchPoints)
	structs.Merge(bp, batchPoints)
	response, err := influx.db.Write(*bp)
	if err != nil || response == nil {
		return nil, err
	}
	if response.Error() != nil {
		return nil, response.Error()
	}
	return &response.Results, nil
}
