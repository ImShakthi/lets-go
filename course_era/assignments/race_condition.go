package main

import "fmt"

/*
------------------------------------------------------------------------------------------------------------------------
What is race Condition?
A race condition occurs when multiple threads try to access and modify the same data (memory address).
E.g., if one thread tries to increase an integer and another thread tries to read it,
this will cause a race condition. On the other hand, there won't be a race condition, if the variable is read-only.
In golang, threads are created implicitly when Goroutines are used.
------------------------------------------------------------------------------------------------------------------------
Why Race condition occurs?

It occurs when two or more process can access and change the shared data at the same time.
It occurred because there were conflicting accesses to a resource .
Critical section problem may cause race condition.
------------------------------------------------------------------------------------------------------------------------
*/
func main() {
	counter := 0
	noOfGoroutines := 10000
	for i := 0; i < noOfGoroutines; i++ {
		go func() {
			// race condition will raise when different go routines, try to
			// access counter at the same time
			counter++
		}()
	}
	fmt.Printf("counter: %v", counter)
}

/*
Execute this file to see race condition: go run -race main.go
--------------------------------------------------------------------
$  go run -race main.go
==================
WARNING: DATA RACE
Read at 0x00c00001c0f8 by goroutine 8:
  main.main.func1()
      /tmp/main.go:10 +0x30

Previous write at 0x00c00001c0f8 by goroutine 7:
  main.main.func1()
      /tmp/main.go:10 +0x44

Goroutine 8 (running) created at:
  main.main()
      /tmp/main.go:9 +0x51

Goroutine 7 (finished) created at:
  main.main()
      /tmp/main.go:9 +0x51
==================
a: 9999
Found 1 data race(s)
exit status 66
--------------------------------------------------------------------
*/
