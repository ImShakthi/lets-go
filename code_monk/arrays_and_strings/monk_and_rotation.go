package arrays_and_strings

import "fmt"

func MonkAndRotation() {
	var t int
	_, _ = fmt.Scanf("%d", &t)

	var arrSize int
	var rotation int
	var temp int
	for i := 0; i < t; i++ {
		_, _ = fmt.Scanf("%d", &arrSize)
		_, _ = fmt.Scanf("%d", &rotation)

		arr := make([]int, 0)
		for ii := 0; ii < arrSize; ii++ {
			_, _ = fmt.Scanf("%d", &temp)
			arr = append(arr, temp)
		}
		printRotatedArr(arr, rotation)
	}
}

func printRotatedArr(input []int, rotation int) {
	n := len(input)
	baseRotIndex := rotation % n

	for i := n - baseRotIndex; i < n; i++ {
		print(input[i], " ")
	}
	for i := 0; i < n-baseRotIndex; i++ {
		print(input[i], " ")
	}
}
