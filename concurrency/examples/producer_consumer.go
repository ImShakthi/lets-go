package examples

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

// The problem describes two processes, the producer and the consumer, who share a common, fixed-size buffer used as a queue.
// The producer's job is to generate data, put it into the buffer, and start again.
// At the same time, the consumer is consuming the data (i.e., removing it from the buffer), one piece at a time.
// The problem is to make sure that the producer won't try to add data into the buffer if it's full and
// that the consumer won't try to remove data from an empty buffer.
// The solution for the producer is to either go to sleep or discard data if the buffer is full.
// The next time the consumer removes an item from the buffer, it notifies the producer, who starts to fill the buffer again.
// In the same way, the consumer can go to sleep if it finds the buffer empty.
// The next time the producer puts data into the buffer, it wakes up the sleeping consumer.

type Consumer struct {
	msgs *chan int
}

// NewConsumer creates a Consumer
func NewConsumer(msgs *chan int) *Consumer {
	return &Consumer{msgs: msgs}
}

// consume reads the msgs channel
func (c *Consumer) consume() {
	fmt.Println("consume: Started")
	for {
		msg := <-*c.msgs
		fmt.Println("consume: Received:", msg)
	}
}

// Producer definition
type Producer struct {
	msgs *chan int
	done *chan bool
}

// NewProducer creates a Producer
func NewProducer(msgs *chan int, done *chan bool) *Producer {
	return &Producer{msgs: msgs, done: done}
}

// produce creates and sends the message through msgs channel
func (p *Producer) produce(max int) {
	fmt.Println("produce: Started")
	for i := 0; i < max; i++ {
		fmt.Println("produce: Sending ", i)
		*p.msgs <- i
	}
	*p.done <- true // signal when done
	fmt.Println("produce: Done")
}

func ProducerConsumer() {
	// profile flags
	cpuProfile := flag.String("cpuProfile", "", "write cpu profile to `file`")
	memProfile := flag.String("memProfile", "", "write memory profile to `file`")

	// get the maximum number of messages from flags
	max := flag.Int("n", 5, "defines the number of messages")

	flag.Parse()

	// utilize the max num of cores available
	runtime.GOMAXPROCS(runtime.NumCPU())

	// CPU Profile
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	var msgs = make(chan int)  // channel to send messages
	var done = make(chan bool) // channel to control when production is done

	// Start a goroutine for Produce.produce
	go NewProducer(&msgs, &done).produce(*max)

	// Start a goroutine for Consumer.consume
	go NewConsumer(&msgs).consume()

	// Finish the program when the production is done
	<-done

	// Memory Profile
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		defer handleFileClose(f)
	}
}

func handleFileClose(f *os.File) {
	if err := f.Close(); err != nil {
		log.Println("error while closing ", err)
	}
}
