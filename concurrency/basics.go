package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func ChanneledPrinter() {
	c := make(chan string)
	go senderChannel(c)
	go printer(c)
}

func PingPongMessenger() {
	wg := new(sync.WaitGroup)
	c := make(chan string, 5)
	wg.Add(1)
	go messengerPing(c, wg)
	wg.Add(1)
	go messengerPong(c, wg)
	go printer(c)
	wg.Wait()
}

func TikTok() {
	tick := time.Tick(time.Second)
	boom := time.After(time.Second * 5)
	for {
		select {
		case <-tick:
			fmt.Println("\ntick")
		case <-boom:
			fmt.Println("\nboom!")
			return
		default:
			fmt.Print(".")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func ChannelFibonacci() {
	fmt.Println(" Fibonacci using channel ")
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

func ImplementChannel() {
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

func ImplementSelect() {

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

func senderChannel(c chan<- string) {
	for i := 0; ; i++ {
		c <- "directed"
	}
}

func messengerPing(c chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		c <- "ping"
	}
	wg.Done()
}

func messengerPong(c chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		c <- "pong"
	}
	wg.Done()
}

func printer(c chan string) {
	for {
		fmt.Println(<-c)
	}
}
