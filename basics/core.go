package basics

import (
	"bytes"
	"fmt"
)

func InitCorePkg() {
	fmt.Println(">>>> CORE PACKAGES <<<")
	i, e := ioFunc("test")
	if e != nil {
		panic("Aiyo rama!!!")
	}
	fmt.Println(">>> i=", i)
}

func ioFunc(msg string) (int, error) {
	var buffer bytes.Buffer
	return buffer.Write([]byte(msg))
}
