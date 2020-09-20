package model

import ()

// Redis configuration
type Redis struct {
	Addr     string `yaml:"addr,omitempty"`
	Password string `yaml:"password,omitempty"`
}
