package arrays_and_strings

import (
	"testing"
)

func TestRotate(t *testing.T) {
	//expected := []int{4, 5, 1, 2, 3}

	input := []int{1, 2, 3, 4, 5}
	printRotatedArr(input, 2)
}
