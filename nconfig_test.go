package nconfig

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig("default")

	hostname := "localhost:8080"
	if hostname != config.Get("hostname") {
		t.Errorf("got %v\nwant %v", config.Get("hostname"), hostname)
	}

	dbuser := "username"
	if dbuser != config.Get("database.user") {
		t.Errorf("got %v\nwant %v", config.Get("database.user"), dbuser)
	}

	dbname := "dbname-dev"
	if dbname != config.Get("database.dbname") {
		t.Errorf("got %v\nwant %v", config.Get("database.dbname"), dbname)
	}

	// t.Log(config)
}

func TestNewConfigSpecifyConfig(t *testing.T) {
	config := NewConfig("production")

	hostname := "example.com"
	if hostname != config.Get("hostname") {
		t.Errorf("got %v\nwant %v", config.Get("hostname"), hostname)
	}

	dbuser := "username"
	if dbuser != config.Get("database.user") {
		t.Errorf("got %v\nwant %v", config.Get("database.user"), dbuser)
	}

	dbname := "dbname-prod"
	if dbname != config.Get("database.dbname") {
		t.Errorf("got %v\nwant %v", config.Get("database.dbname"), dbname)
	}

	// t.Log(config)
}
