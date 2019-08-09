package main

import (
	"fmt"
	"lets-go/basics"
	"lets-go/concurrency"
)

func main() {
	fmt.Println("Lets go...")
	concurrency.Init()
	basics.Init()
	basics.InitCorePkg()
}
