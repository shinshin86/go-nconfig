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
