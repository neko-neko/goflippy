package main

import "github.com/kelseyhightower/envconfig"

// Specification variables on which goflippy depends
type Specification struct {
	Debug        bool     `default:"false" envconfig:"DEBUG"`
	AllowOrigins []string `default:"*" envconfig:"ALLOW_ORIGINS"`

	// For Mongo settings
	StoreTimeoutSeconds int      `default:"30" envconfig:"STORE_TIMEOUT_SECONDS"`
	StoreAddrs          []string `default:"mongo" envconfig:"STORE_ADDRS"`
	StoreDB             string   `default:"goflippy" envconfig:"STORE_DB"`
	StoreUser           string   `default:"" envconfig:"STORE_USER"`
	StorePassword       string   `default:"" envconfig:"STORE_PASSWORD"`
	StoreSource         string   `default:"" envconfig:"STORE_SOURCE"`
}

// Spec is global env instance
var Spec Specification

// EnvInit environment variables
func EnvInit() error {
	return envconfig.Process("", &Spec)
}
