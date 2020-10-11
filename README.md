# go-nconfig

[![Build Status](https://travis-ci.org/shinshin86/go-nconfig.svg?branch=master)](https://travis-ci.org/shinshin86/go-nconfig)

Like a npm module of [node-config](<https://github.com/lorenwest/node-config>)  a go package. 



## Quick Start

Install in your app directory. And edit the default config file.

```bash
go mod init github.com/you/hello
go get github.com/shinshin86/go-nconfig
mkdir config
vim config/default.json
```



Setup `default.json`

```bash
{
  "hostname": "localhost:8080",
  "database": {
    "protocol": "tcp(127.0.0.1:3306)",
    "dbname": "dbname-dev"
  }
}
```



Create `main.go`.

```bash
vim main.go
```



```go
package main

import (
	"fmt"

	"github.com/shinshin86/go-nconfig"
)

func main() {
	config := nconfig.New("default")

	fmt.Println("hostname : ", config.Get("hostname"))
	fmt.Println("database.protocol : ", config.Get("database.protocol"))
	fmt.Println("database.dbname : ", config.Get("database.dbname"))
}
```



```bash
go run main.go
# ==> hostname :  localhost:8080
# ==> database.protocol :  tcp(127.0.0.1:3306)
# ==> database.dbname :  dbname-dev
```



##  When specify config file

Reading specify config file.
( Not set value is refer to default.json )



default.json

```json
{
  "hostname": "localhost:8080",
  "database": {
    "protocol": "tcp(127.0.0.1:3306)",
    "dbname": "dbname-dev"
  },
  "dep1": {
    "val": "dev1",
    "dep2": {
      "val": "dev2",
      "dep3": {
        "val": "dev3",
        "dep4": {
          "val": "dev4",
          "dep5": {
            "val": "dev5",
            "dep6": {
              "val": "dev6"
            }
          }
        }
      }
    }
  }
}

```



production.json

```json
{
  "hostname": "example.com",
  "database": {
    "dbname": "dbname-prod"
  },
  "dep1": {
    "val": "prod1",
    "dep2": {
      "val": "prod2",
      "dep3": {
        "val": "prod3",
        "dep4": {
          "val": "prod4",
          "dep5": {
            "val": "prod5",
            "dep6": {
              "val": "prod6"
            }
          }
        }
      }
    }
  }
}

```



main.go

```go
package main

import (
	"fmt"

	"github.com/shinshin86/go-nconfig"
)

func main() {
	config := nconfig.New("production")

	fmt.Println("hostname : ", config.Get("hostname"))
	fmt.Println("database.protocol : ", config.Get("database.protocol"))
	fmt.Println("database.dbname : ", config.Get("database.dbname"))
	fmt.Println("dep6 value : ", config.Get("dep1.dep2.dep3.dep4.dep5.dep6.val"))
}
```



```bash
go run main.go
# ==> hostname :  example.com
# ==> database.protocol :  tcp(127.0.0.1:3306)
# ==> database.dbname :  dbname-prod
# ==> dep6 value :  prod6
```