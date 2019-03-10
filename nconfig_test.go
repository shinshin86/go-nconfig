package nconfig

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config := LoadConfig("default")

	dbuser := "username"
	configDbUser := config["database"].(map[string]interface{})["user"]
	if dbuser != configDbUser {
		t.Errorf("got %v\nwant %v", configDbUser, dbuser)
	}
	dbname := "dbname-dev"
	configDbname := config["database"].(map[string]interface{})["dbname"]
	if dbname != configDbname {
		t.Errorf("got %v\nwant %v", configDbname, dbname)
	}

	hostname := "localhost:8080"
	configHostname := config["hostname"]

	if hostname != configHostname {
		t.Errorf("got %v\nwant %v", configHostname, hostname)
	}
	// t.Log(config)
}

func TestLoadConfigSpecifyConfig(t *testing.T) {
	config := LoadConfig("production")

	dbuser := "username"
	configDbUser := config["database"].(map[string]interface{})["user"]
	if dbuser != configDbUser {
		t.Errorf("got %v\nwant %v", configDbUser, dbuser)
	}
	dbname := "dbname-prod"
	configDbname := config["database"].(map[string]interface{})["dbname"]
	if dbname != configDbname {
		t.Errorf("got %v\nwant %v", configDbname, dbname)
	}

	hostname := "example.com"
	configHostname := config["hostname"]

	if hostname != configHostname {
		t.Errorf("got %v\nwant %v", configHostname, hostname)
	}
	// t.Log(config)
}
