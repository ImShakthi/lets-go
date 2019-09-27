package main

import (
	"fmt"
	"lets-go/snippets"
)

func main() {
	fmt.Println("Lets go...")
	//concurrency.Init()
	//basics.Init()
	//basics.InitCorePkg()
	//snippets.NewTomlParser()

	errorHandler := snippets.NewErrorHandler()
	errorHandler.Play()
}
