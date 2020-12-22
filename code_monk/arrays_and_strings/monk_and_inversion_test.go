package arrays_and_strings

import (
	"fmt"
	"testing"
)

var input = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

func TestMatrix_Start(t *testing.T) {
	Start()
}

func TestMatrix_Print(t *testing.T) {
	m := matrix{
		Data:       input,
		NofRows:    3,
		NofColumns: 3,
	}

	m.Print()
}

func TestMatrix_GetInversion(t *testing.T) {
	//m := matrix{
	//	Data:       input,
	//	NofRows:    3,
	//	NofColumns: 3,
	//}

	m := matrix{
		Data:       [][]int{{4, 3}, {1, 4}},
		NofRows:    2,
		NofColumns: 2,
	}

	fmt.Println(">>>>Inversion= ", m.GetInversion())
}
