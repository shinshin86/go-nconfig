# go-nconfig (WIP)

[![Build Status](https://travis-ci.org/shinshin86/go-nconfig.svg?branch=master)](https://travis-ci.org/shinshin86/go-nconfig)

Like a npm module of config  a go package. 



## Usage

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

        fmt.Println("Not supported read of six level : ", config.Get("dep1.dep2.dep3.dep4.dep5.dep6.val"))
}
```



#### Notes :

Depth supports to five levels.

