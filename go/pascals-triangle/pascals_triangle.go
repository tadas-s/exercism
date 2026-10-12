package pascalstriangle

func Triangle(n int) [][]int {
	allocate := max(0, n)

	if allocate == 0 {
		return [][]int{}
	}

	triangle := make([][]int, allocate)
	triangle[0] = make([]int, 1)
	triangle[0][0] = 1

	for row := 1; row < n; row++ {
		triangle[row] = make([]int, row+1)

		for column := range row + 1 {
			a := 0

			if column > 0 {
				a = triangle[row-1][column-1]
			}

			b := 0

			if column < row {
				b = triangle[row-1][column]
			}

			triangle[row][column] = a + b
		}
	}

	return triangle
}
