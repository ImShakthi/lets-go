package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func Init() {
	fmt.Println(">>> Concurrency <<<")

	//testFibo()
	//
	//c := make(chan string)
	//
	//go messengerPing(c)
	//go messengerPong(c)
	//go senderChannel(c)
	//go printer(c, 1)
	//go printer(c, 2)
	//go printer2(c)
	//
	//testSelect()
	//testChannel()

	//tikTok()

	test()

	var input string
	fmt.Scanln(&input)

	fmt.Print(input)
}

func test() {
	c := make(chan int)
	go input(c)
	go print(c)
}

func input(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
}

func print(c chan int) {
	for {
		fmt.Println("-> ", <-c)
		time.Sleep(time.Second)
	}
}

func tikTok() {
	tick := time.Tick(time.Second)
	boom := time.After(time.Second * 5)
	for {
		select {
		case <-tick:
			fmt.Print("tick")
		case <-boom:
			fmt.Println("\nboom!")
			return
		default:
			fmt.Print(".")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func testFibo() {
	fmt.Println(" TEST FIBO ")
	n := make(chan int, 2)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-n)
		}
		quit <- true
	}()
	fibo(n, quit)
}

func fibo(n chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case n <- x:
			//fmt.Printf("\nx=%v", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
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
