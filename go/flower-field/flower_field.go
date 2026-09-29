package flowerfield

import (
	"errors"
	"strconv"
	"strings"
)

var directions = [][2]int{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0} /*     */, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

// Annotate returns an annotated board
func Annotate(board []string) []string {
	var result []string

	for y := range len(board) {
		var row strings.Builder

		for x := range len(board[0]) {
			flowers := 0
			thisSquare, _ := squareAt(board, x, y)

			if thisSquare == '*' {
				row.WriteString(string(thisSquare))
				continue
			}

			for _, dir := range directions {
				b, err := squareAt(board, x+dir[0], y+dir[1])

				if err == nil && b == '*' {
					flowers++
				}
			}

			if flowers > 0 {
				row.WriteString(strconv.Itoa(flowers))
			} else {
				row.WriteString(string(thisSquare))
			}
		}

		result = append(result, row.String())
	}

	return result
}

func squareAt(board []string, x, y int) (byte, error) {
	if y >= 0 && y < len(board) && x >= 0 && x < len(board[0]) {
		return board[y][x], nil
	}

	return ' ', errors.New("out of bounds")
}
