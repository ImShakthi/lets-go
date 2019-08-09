package basics

import (
	"bytes"
	"fmt"
)

type Student struct {
	Name    string `json:"name" binding:"required"`
	Checked bool   `json:"checked" binding:"required"`
}

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
