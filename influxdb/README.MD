# Influxdb CLient

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://github.com/go-iot-platform/influxdb)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/go-iot-platform/influxdb)

Merge two different structs, and overwrite fields

## What is this

influxdb client

## Getting Started!
> model
```
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
```
> How to use
  ```
   con, err := influxdb.NewClient(config)
   con.Query(query)
   con.Write(batchPoints)
  ```


### Installation

Request handler requires [Go](https://golang.org/) v1.11+ to run.

Install the package.

```sh
$ go get github.com/go-iot-platform/influxdb
```

#### Kubernetes + Google Cloud

See [KUBERNETES.md](https://github.com/joemccann/dillinger/blob/master/KUBERNETES.md)


### Todos

 - Write MORE Tests
 - Add Night Mode

License
----

MIT
