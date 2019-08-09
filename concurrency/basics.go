package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func Init() {
	fmt.Print(">>> Concurrency <<<")

	c := make(chan string)

	go messengerPing(c)
	go messengerPong(c)
	go senderChannel(c)
	go printer(c, 1)
	go printer(c, 2)
	go printer2(c)

	testSelect()

	testChannel()

	var input string
	fmt.Scanln(&input)

	fmt.Print(input)
}

func testChannel() {
	done := make(chan bool)
	fmt.Println("Start go routine")
	go sayHello(done)
	fmt.Println("Did say hello run ? ", <-done)
}

func sayHello(done chan bool) {
	fmt.Println("In say hello go routine")
	time.Sleep(time.Second * 10)
	done <- true
}

func testSelect() {

	c1 := make(chan string, 3)
	c2 := make(chan string, 5)

	go func() {
		for {
			c1 <- " from C1"
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			c2 <- " from C2"
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			select {
			case msg := <-c1:
				fmt.Println(msg)
			case msg := <-c2:
				fmt.Println(msg)
			case <-time.After(time.Second):
				fmt.Println("time out")
				//default:
				//	fmt.Println("nothing is ready")
			}
		}
	}()
}

func namePrinter(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, "->", name, ", ")

		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func senderChannel(c chan<- string) {
	for i := 0; ; i++ {
		c <- "directed"
	}
}

func messengerPing(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func messengerPong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string, id int) {
	for {
		fmt.Println("id=", id, ", ", <-c)
		time.Sleep(time.Second * 1)
	}
}

func printer2(c chan string) {
	for {
		fmt.Println("second printer=", <-c)
		time.Sleep(time.Second * 1)
	}
}
