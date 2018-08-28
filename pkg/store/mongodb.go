package store

import (
	"time"

	"gopkg.in/mgo.v2"
)

// singleton instance
var store *DB

// Configuration is MongoDB configuration
type Configuration struct {
	Addrs          []string
	TimeoutSeconds int
	DB             string
	User           string
	Password       string
	Source         string
}

// DB datastore
type DB struct {
	conf    Configuration
	Session *mgo.Session
	MongoDB *mgo.Database
}

// Init return
func Init(conf Configuration) *DB {
	store = &DB{
		conf: conf,
	}

	return store
}

// Connect MongoDB
func (d *DB) Connect() error {
	dial := &mgo.DialInfo{
		Addrs:    d.conf.Addrs,
		Timeout:  time.Duration(d.conf.TimeoutSeconds) * time.Second,
		Database: d.conf.DB,
		Username: d.conf.User,
		Password: d.conf.Password,
		Source:   d.conf.Source,
	}
	session, err := mgo.DialWithInfo(dial)
	if err != nil {
		return err
	}

	d.Session = session
	d.MongoDB = session.DB("")

	return nil
}

// Close connection
func (d *DB) Close() error {
	d.Session.Close()

	return nil
}
