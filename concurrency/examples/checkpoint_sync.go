package examples

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// The checkpoint synchronization is a problem of synchronizing multiple tasks.
// Consider a workshop where several workers assembling details of some mechanism.
// When each of them completes his work, they put the details together.
// There is no store, so a worker who finished its part first must wait for others before starting another one.
// Putting details together is the checkpoint at which tasks synchronize themselves before going their paths apart.

var (
	partList    = []string{"A", "B", "C", "D"}
	nAssemblies = 3
)

func worker(part string, wg *sync.WaitGroup) {
	log.Println(part, "worker begins part")
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	log.Println(part, "worker completes part")
	wg.Done()
}

func CheckpointSync() {
	rand.Seed(time.Now().UnixNano())
	wg := new(sync.WaitGroup)
	for c := 1; c <= nAssemblies; c++ {
		log.Println("begin assembly cycle", c)
		wg.Add(len(partList))
		for _, part := range partList {
			go worker(part, wg)
		}
		wg.Wait()
		log.Println("assemble.  cycle", c, "complete")
	}
}
