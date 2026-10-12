package rectangles

func Count(diagram []string) int {
	rows := len(diagram)

	if rows == 0 {
		return 0
	}

	columns := len(diagram[0])

	if columns == 0 {
		return 0
	}

	rectangles := 0

	for r1 := 0; r1 < rows-1; r1++ {
		for c1 := 0; c1 < columns-1; c1++ {
			for r2 := r1 + 1; r2 < rows; r2++ {
				for c2 := c1 + 1; c2 < columns; c2++ {
					if isValidRectangle(diagram, r1, c1, r2, c2) {
						rectangles++
					}
				}
			}
		}
	}

	return rectangles
}

func isValidRectangle(diagram []string, r1, c1, r2, c2 int) bool {
	if diagram[r1][c1] != '+' || diagram[r1][c2] != '+' ||
		diagram[r2][c1] != '+' || diagram[r2][c2] != '+' {
		return false
	}

	for r := r1 + 1; r < r2; r++ {
		if diagram[r][c1] != '|' && diagram[r][c1] != '+' {
			return false
		}

		if diagram[r][c2] != '|' && diagram[r][c2] != '+' {
			return false
		}
	}

	for c := c1 + 1; c < c2; c++ {
		if diagram[r1][c] != '-' && diagram[r1][c] != '+' {
			return false
		}

		if diagram[r2][c] != '-' && diagram[r2][c] != '+' {
			return false
		}
	}

	return true
}
