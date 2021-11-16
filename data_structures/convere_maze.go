package data_structures

//type node struct {
//	Index int
//	Next  *node
//}

func ConstructMaze() {
	input := []int{4, 4, 1, 4, 13, 8, 8, 8, 0, 8, 14, 9, 15, 11, -1, 10, 15, 22, 22, 22, 22, 22, 21}

	visited := make([]bool, 23)

	src := 9
	dst := 2

	i := src
	for {
		if visited[input[i]]{
			break
		}
		visited[input[i]] = true
		i = input[i]
	}

	i = dst
	for {
		if visited[input[i]]{
			println(">>>>NEAREST::", input[i])
			break
		}
		visited[input[i]] = true
		i = input[i]
	}
}
