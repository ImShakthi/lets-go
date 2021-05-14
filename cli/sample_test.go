package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	if !isEven(10) {
		t.Errorf("not an even number")
	}
}
