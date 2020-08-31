# go-nconfig (WIP)

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
    "user": "username",
    "pass": "password",
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
	fmt.Println("dbname : ", config.Get("database.user"))
}
```



```bash
go run main.go
# ==> hostname :  localhost:8080
# ==> dbname :  username
```



##  When specify config file

Reading specify config file.
( Not set value is refer to default.json )

```go
package main

import (
        "fmt"

        "github.com/shinshin86/go-nconfig"
)

func main() {
        config := nconfig.New("production")

        fmt.Println("=== Read config sample ===")
        fmt.Println("hostname : ", config.Get("hostname"))
        fmt.Println("dbname : ", config.Get("database.user"))
        fmt.Println("dep5 value : ", config.Get("dep1.dep2.dep3.dep4.dep5.val"))
}
```

