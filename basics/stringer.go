package basics

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s is %d with pointer", p.Name, p.Age)
}

type Person2 struct {
	Name string
	Age  int
}

func (p Person2) String() string {
	return fmt.Sprintf("%s is %d without pointer", p.Name, p.Age)
}

func (p Person2) GoString() string {
	return fmt.Sprintf("{Wooow da!}")
}

type Person3 struct {
	Name string
	Age  int
}

func (p *Person3) String() string {
	return fmt.Sprintf("%s is %d with pointer with value call", p.Name, p.Age)
}
