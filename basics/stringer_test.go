package basics_test

import (
	"fmt"
	"lets-go/basics"
	"testing"
)

func TestStringerWithPointer(t *testing.T) {
	p := &basics.Person{
		Name: "Shakthi",
		Age:  25,
	}
	fmt.Println(p)
}

func TestStringerWithoutPointer(t *testing.T) {
	p := basics.Person2{
		Name: "Shakthi",
		Age:  25,
	}
	fmt.Println(p)
	fmt.Printf("%#v", p)
}

func TestStringerWithPointerCallAsValue(t *testing.T) {
	p := basics.Person3{
		Name: "Shakthi",
		Age:  25,
	}
	fmt.Println(p)
}

func TestInterfaces(t *testing.T) {
	var p *basics.Person
	var p1 *basics.Person
	var s fmt.Stringer
	fmt.Println(p == p1)
	fmt.Printf("%#v \n", p)
	fmt.Printf("%#v", s)
}

func TestCornerCaseLikeShift(t *testing.T) {
	fmt.Printf("%T %T", 2.0, 2.0<<0)
}
