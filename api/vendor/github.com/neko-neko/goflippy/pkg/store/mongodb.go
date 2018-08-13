package store

import (
	"gopkg.in/mgo.v2"
)

// singleton instance
var store *DB

// Configuration is MongoDB configuration
type Configuration struct {
	URL string
	DB  string
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
	session, err := mgo.Dial(d.conf.URL)
	if err != nil {
		return err
	}

	d.Session = session
	d.MongoDB = session.DB(d.conf.DB)

	return nil
}

// Close connection
func (d *DB) Close() error {
	d.Session.Close()

	return nil
}
