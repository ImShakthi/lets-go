package examples

import (
	"fmt"
	"sync"
	"time"
)

// The barber has one barber's chair in a cutting room and a waiting room containing a number of chairs in it.
// When the barber finishes cutting a customer's hair, he dismisses the customer and goes to the waiting room
// to see if there are others waiting. If there are, he brings one of them back to the chair and cuts their hair.
// If there are none, he returns to the chair and sleeps in it. Each customer, when they arrive, looks to see what
// the barber is doing. If the barber is sleeping, the customer wakes him up and sits in the cutting room chair.
// If the barber is cutting hair, the customer stays in the waiting room. If there is a free chair in the waiting
// room, the customer sits in it and waits their turn. If there is no free chair, the customer leaves.

const (
	sleeping = iota
	checking
	cutting
)

var stateLog = map[int]string{
	0: "Sleeping",
	1: "Checking",
	2: "Cutting",
}

type Barber struct {
	name string
	sync.Mutex
	state    int // Sleeping/Checking/Cutting
	customer *Customer
}

type Customer struct {
	name string
}

func (c *Customer) String() string {
	return fmt.Sprintf("%p", c)[7:]
}

func NewBarber() (b *Barber) {
	return &Barber{
		name:  "Sam",
		state: sleeping,
	}
}

// Barber goroutine
// Checks for customers
// Sleeps - wait for wakers to wake him up
func barber(b *Barber, wr chan *Customer, wakers chan *Customer, wg *sync.WaitGroup) {
	for {
		b.Lock()
		defer b.Unlock()
		b.state = checking
		b.customer = nil

		// checking the waiting room
		fmt.Printf("Checking waiting room: %d\n", len(wr))
		time.Sleep(time.Millisecond * 100)
		select {
		case c := <-wr:
			HairCut(c, b, wg)
			b.Unlock()
		default: // Waiting room is empty
			fmt.Printf("Sleeping Barber - %s\n", b.customer)
			b.state = sleeping
			b.customer = nil
			b.Unlock()
			c := <-wakers
			b.Lock()
			fmt.Printf("Woken by %s\n", c)
			HairCut(c, b, wg)
			b.Unlock()
		}
	}
}

func HairCut(c *Customer, b *Barber, wg *sync.WaitGroup) {
	b.state = cutting
	b.customer = c
	b.Unlock()
	fmt.Printf("Cutting  %s hair\n", c)
	time.Sleep(time.Millisecond * 100)
	b.Lock()
	wg.Done()
	b.customer = nil
}

// customer goroutine
// just fizzles out if it's full, otherwise the customer
// is passed along to the channel handling it's haircut etc
func customer(c *Customer, b *Barber, wr chan<- *Customer, wakers chan<- *Customer, wg *sync.WaitGroup) {
	// arrive
	time.Sleep(time.Millisecond * 50)
	// Check on barber
	b.Lock()
	fmt.Printf("Customer %s checks %s barber | room: %d, w %d - customer: %s\n",
		c, stateLog[b.state], len(wr), len(wakers), b.customer)
	switch b.state {
	case sleeping:
		select {
		case wakers <- c:
		default:
			select {
			case wr <- c:
			default:
				wg.Done()
			}
		}
	case cutting:
		select {
		case wr <- c:
		default: // Full waiting room, leave shop
			wg.Done()
		}
	case checking:
		panic("Customer shouldn't check for the Barber when Barber is Checking the waiting room")
	}
	b.Unlock()
}

func SleepingBarber() {
	b := NewBarber()
	b.name = "Rocky"
	WaitingRoom := make(chan *Customer, 5) // 5 chairs
	wakers := make(chan *Customer, 1)      // Only one waker at a time
	wg := new(sync.WaitGroup)
	go barber(b, WaitingRoom, wakers, wg)

	time.Sleep(time.Millisecond * 100)
	n := 10
	wg.Add(10)
	// Spawn customers
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond * 50)
		c := new(Customer)
		go customer(c, b, WaitingRoom, wakers, wg)
	}

	wg.Wait()
	fmt.Println("No more customers for the day")
}
