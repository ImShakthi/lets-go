package basics

import (
	"fmt"
	"os"
)

func Init() {
	fmt.Println(">>>>> FUNCTIONS <<<<<")
	fmt.Println("Sum= ", Add(1, 2, 3, 4, 5))
	fmt.Println("Factorial of 5=", Factorial(5))
	closure()
	checkDefer()
	fileManager()
	panicRecover()
}

func Add(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// closure
func closure() {
	fmt.Println(">>> CLOSURE <<< ")
	x := 0
	increment := func() int {
		x += 5
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

// defer
func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func checkDefer() {
	defer second()
	first()
}

// panic and recover
func panicRecover() {
	fmt.Println(">>>>>>>>>> PANIC AND RECOVER <<<<<<<<<<<")
	panic("AIYOOOOO RAMA PANIC DA!!!!")
	//recoverMsg := recover()
	//fmt.Println(recoverMsg)
}

func fileManager() {
	file, err := os.Open("/Users/sakthivel/go/src/lets-go/basics/text.txt")
	if err != nil {
		fmt.Println("error while opening file, ", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		fmt.Println("error while getting file stat, ", err)
		return
	}
	fmt.Println("file name is ", info.Name(), ", ", info.ModTime())
	writeString, err := file.WriteString("append this data")
	if err != nil {
		fmt.Println(" error while appending data, ", err)
	}
	fmt.Println("writeString=", writeString)
}
