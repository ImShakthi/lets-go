package awesome

import "fmt"

// User defined function type
type math func(a int, b int) int

func Kand() {
	// Anonymous function
	anonFn := func() {
		fmt.Println("Anonymous function")
	}
	anonFn()
	fmt.Printf("%T", anonFn)

	func() {
		fmt.Println("Anonymous function - 2")
	}()

	var add math = func(a int, b int) int {
		return a + b
	}
	fmt.Println("sum=", add(5, 6))

	var prod math = func(a int, b int) int {
		return a * b
	}
	fmt.Println("prod=", prod(5, 6))

	userDefinedFuncHandler(10, 11, func(a int, b int) int {
		return a - b
	})

}

// High order function
func userDefinedFuncHandler(a int, b int, fn math) int {
	result := fn(a, b)
	fmt.Println("result=", result)
	return result
}
