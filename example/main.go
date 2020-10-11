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
