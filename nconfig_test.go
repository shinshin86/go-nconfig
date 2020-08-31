package nconfig

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	configname := "default"
	config := New(configname)

	if configname != config.Name {
		t.Errorf("got %v\nwant %v", config.Name, configname)
	}

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

	apiakey := "foobar"
	if apiakey != config.Get("apis.servicea.key") {
		t.Errorf("got %v\nwant %v", config.Get("apis.servicea.key"), apiakey)
	}

	apibkey := "barfoo"
	if apibkey != config.Get("apis.serviceb.key") {
		t.Errorf("got %v\nwant %v", config.Get("apis.serviceb.key"), apibkey)
	}

	// deptch check
	dep1 := "dev1"
	if dep1 != config.Get("dep1.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.val"), dep1)
	}
	dep2 := "dev2"
	if dep2 != config.Get("dep1.dep2.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.val"), dep2)
	}
	dep3 := "dev3"
	if dep3 != config.Get("dep1.dep2.dep3.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.dep3.val"), dep3)
	}
	dep4 := "dev4"
	if dep4 != config.Get("dep1.dep2.dep3.dep4.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.dep3.dep4.val"), dep4)
	}
	dep5 := "dev5"
	if dep5 != config.Get("dep1.dep2.dep3.dep4.dep5.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.dep3.dep4.dep5.val"), dep5)
	}
	dep6 := "dev6"
	if dep6 != config.Get("dep1.dep2.dep3.dep4.dep5.dep6.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.dep3.dep4.dep5.dep6.val"), dep6)
	}
	// t.Log(config)
}

func TestNewConfigSpecifyConfig(t *testing.T) {
	configname := "production"
	config := New(configname)

	if configname != config.Name {
		t.Errorf("got %v\nwant %v", config.Name, configname)
	}

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

	apiakey := "keyaaa"
	if apiakey != config.Get("apis.servicea.key") {
		t.Errorf("got %v\nwant %v", config.Get("apis.servicea.key"), apiakey)
	}

	apibkey := "barfoo"
	if apibkey != config.Get("apis.serviceb.key") {
		t.Errorf("got %v\nwant %v", config.Get("apis.serviceb.key"), apibkey)
	}

	// deptch check
	dep1 := "prod1"
	if dep1 != config.Get("dep1.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.val"), dep1)
	}
	dep2 := "prod2"
	if dep2 != config.Get("dep1.dep2.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.val"), dep2)
	}
	dep3 := "prod3"
	if dep3 != config.Get("dep1.dep2.dep3.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.dep3.val"), dep3)
	}
	dep4 := "prod4"
	if dep4 != config.Get("dep1.dep2.dep3.dep4.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.dep3.dep4.val"), dep4)
	}
	dep5 := "prod5"
	if dep5 != config.Get("dep1.dep2.dep3.dep4.dep5.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.dep3.dep4.dep5.val"), dep5)
	}
	dep6 := "prod6"
	if dep6 != config.Get("dep1.dep2.dep3.dep4.dep5.dep6.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.dep3.dep4.dep5.dep6.val"), dep6)
	}
	// t.Log(config)
}

func TestNewConfigNoNamesSpecified(t *testing.T) {
	configname := "default"
	config := New("")

	if configname != config.Name {
		t.Errorf("got %v\nwant %v", config.Name, configname)
	}

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

	apiakey := "foobar"
	if apiakey != config.Get("apis.servicea.key") {
		t.Errorf("got %v\nwant %v", config.Get("apis.servicea.key"), apiakey)
	}

	apibkey := "barfoo"
	if apibkey != config.Get("apis.serviceb.key") {
		t.Errorf("got %v\nwant %v", config.Get("apis.serviceb.key"), apibkey)
	}

	// deptch check
	dep1 := "dev1"
	if dep1 != config.Get("dep1.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.val"), dep1)
	}
	dep2 := "dev2"
	if dep2 != config.Get("dep1.dep2.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.val"), dep2)
	}
	dep3 := "dev3"
	if dep3 != config.Get("dep1.dep2.dep3.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.dep3.val"), dep3)
	}
	dep4 := "dev4"
	if dep4 != config.Get("dep1.dep2.dep3.dep4.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.dep3.dep4.val"), dep4)
	}
	dep5 := "dev5"
	if dep5 != config.Get("dep1.dep2.dep3.dep4.dep5.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.dep3.dep4.dep5.val"), dep5)
	}
	dep6 := "dev6"
	if dep6 != config.Get("dep1.dep2.dep3.dep4.dep5.dep6.val") {
		t.Errorf("got %v\nwant %v", config.Get("dep1.dep2.dep3.dep4.dep5.dep6.val"), dep6)
	}
	// t.Log(config)
}
