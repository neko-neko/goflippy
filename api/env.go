package main

import "github.com/kelseyhightower/envconfig"

// Specification variables on which goflippy depends
type Specification struct {
	Debug bool `default:"false"`
	Port  int  `default:"8080"`

	// For Mongo settings
	StoreURL string `default:"mongodb://mongo"`
	DB       string `default:"goflippy"`
}

// Spec is global env instance
var Spec Specification

// EnvInit environment variables
func EnvInit() error {
	return envconfig.Process("", &Spec)
}
