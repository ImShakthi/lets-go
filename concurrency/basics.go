package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func Init() {
	fmt.Print(">>> Concurrency <<<")

	var c chan string = make(chan string)

	go pinger(c)
	go printers(c)

	var input string
	fmt.Scanln(&input)

	fmt.Print(input)
}

func printer(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, "->", name, ", ")

		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printers(c chan string) {
	for {
		msg := <- c
		fmt.Println(">>>>" ,msg)
		time.Sleep(time.Second * 1)
	}
}