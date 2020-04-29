package unicodes_test

import (
	"lets-go/unicodes"
	"testing"
)

func Testவersion(t *testing.T) {
	unicodes.Iவersion()
}

func TestEmptyStruct(t *testing.T){
	var e *unicodes.Empty
	e.Print()
}
