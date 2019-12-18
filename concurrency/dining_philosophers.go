package concurrency

import (
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Number of philosophers is simply the length of this list.
var ph = []string{"Mark", "Russell", "Rocky", "Haris", "Root"}

const hunger = 3                // Number of times each philosopher eats
const think = time.Second / 100 // Mean think time
const eat = time.Second / 100   // Mean eat time

var logger = log.New(os.Stdout, "", 0)

var dining sync.WaitGroup

func diningProblem(phName string, dominantHand, otherHand *sync.Mutex) {
	logger.Println(phName, "Seated")
	h := fnv.New64a()
	_, err := h.Write([]byte(phName))
	if err != nil {
		log.Println(err)
		return
	}

	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := hunger; h > 0; h-- {
		logger.Println(phName, "Hungry")
		dominantHand.Lock() // pick up forks
		otherHand.Lock()
		logger.Println(phName, "Eating")
		rSleep(eat)
		dominantHand.Unlock() // put down forks
		otherHand.Unlock()
		logger.Println(phName, "Thinking")
		rSleep(think)
	}
	logger.Println(phName, "Satisfied")
	dining.Done()
	logger.Println(phName, "Left the table")
}

func DiningPhilosophersProblem() {
	logger.Println("Table empty")
	fork0 := &sync.Mutex{}
	forkLeft := fork0
	for i := 1; i < len(ph); i++ {
		dining.Add(1)
		forkRight := &sync.Mutex{}
		go diningProblem(ph[i], forkLeft, forkRight)
		forkLeft = forkRight
	}
	go diningProblem(ph[0], fork0, forkLeft)
	dining.Wait() // wait for philosphers to finish
	logger.Println("Table empty")
}
