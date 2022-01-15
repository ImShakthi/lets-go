package concurrency

import (
	"fmt"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func checkPanicRecover() {
	f1()
}
func f1() {
	defer func() {
		fmt.Println("Defer in f1")
		if r := recover(); r != nil {
			fmt.Println("RECOVERED IN F1")
		}
	}()
	f2()
	fmt.Println("After panic in f1")
}
func f2() {

	func() {
		defer func() {
			fmt.Println("Defer in f2")
			if r := recover(); r != nil {
				fmt.Println("RECOVERED IN F2")

				if err, ok := r.(stackTracer); ok {
					fmt.Printf(">>>> %+v ", err)
				}else {
					fmt.Println("NOOOOOOO!")
				}

			}
		}()

		panic("Panic Demo")
	}()

	fmt.Println(">>>HOLA!!")

}
