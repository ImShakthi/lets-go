package arrays_and_strings

import "fmt"

type matrix struct {
	Data         [][]int
	NofRows      int
	NofColumns   int
	NofInversion int
}

func Start() {
	fmt.Printf("Monk and inversion\n")

	var n int
	_, _ = fmt.Scanf("%d", &n)

	result := make([]int, 0)

	for i := 0; i < n; i++ {
		var matrixSize int
		_, _ = fmt.Scanf("%d", &matrixSize)

		input := [][]int{}
		var temp int
		for ii := 0; ii < matrixSize; ii++ {
			arr := []int{}
			for jj := 0; jj < matrixSize; jj++ {
				_, _ = fmt.Scanf("%d", &temp)
				arr = append(arr, temp)
			}
			input = append(input, arr)
		}
		m := matrix{
			Data:       input,
			NofRows:    matrixSize,
			NofColumns: matrixSize,
		}
		result = append(result, m.GetInversion())
	}

	for _, r := range result {
		fmt.Println(r)
	}
}
func (m *matrix) GetInversion() int {
	for x1 := 0; x1 < m.NofRows; x1++ {
		for y1 := 0; y1 < m.NofColumns; y1++ {
			m.computeInversion(x1, y1)
		}
	}
	return m.NofInversion
}

func (m *matrix) Print() {
	fmt.Printf("Data in %v x %v matrix is ::\n", m.NofRows, m.NofColumns)
	for i := 0; i < m.NofRows; i++ {
		for j := 0; j < m.NofColumns; j++ {
			fmt.Printf("%v ", m.Data[i][j])
		}
		fmt.Println()
	}
}

func (m *matrix) isInversion(x1, y1, x2, y2 int) bool {
	return m.Data[x1][y1] > m.Data[x2][y2] && x1 <= x2 && y1 <= y2
}

func (m *matrix) computeInversion(x1 int, y1 int) {
	for x2 := 0; x2 < m.NofRows; x2++ {
		for y2 := 0; y2 < m.NofColumns; y2++ {
			if m.isInversion(x1, y1, x2, y2) {
				m.NofInversion = m.NofInversion + 1
				fmt.Printf("\nInversion at x1,y1=%v,%v x2,y2=%v,%v ", x1, y1, x2, y2)
			}
		}
	}
}
