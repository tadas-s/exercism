package saddlepoints

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Matrix [][]int

type Pair [2]int

func New(s string) (*Matrix, error) {
	var points Matrix

	for rowString := range strings.Lines(s) {
		var row []int

		for pointString := range strings.FieldsSeq(rowString) {
			if point, err := strconv.Atoi(pointString); err == nil {
				row = append(row, point)
			} else {
				return nil, fmt.Errorf("cannot parse matrix string: [%w]", err)
			}
		}

		points = append(points, row)
	}

	return &points, nil
}

func (m *Matrix) Saddle() []Pair {
	var pairs []Pair

	if len(*m) == 0 {
		return pairs
	}

	rowMaxValues := make([]int, len(*m))

	for i := range len(rowMaxValues) {
		rowMaxValues[i] = math.MinInt
	}

	columnMinValues := make([]int, len((*m)[0]))

	for i := range len(columnMinValues) {
		columnMinValues[i] = math.MaxInt
	}

	for row, rowPoints := range *m {
		for column, point := range rowPoints {
			rowMaxValues[row] = max(rowMaxValues[row], point)
			columnMinValues[column] = min(columnMinValues[column], point)
		}
	}

	for row, rowPoints := range *m {
		for column, point := range rowPoints {
			if rowMaxValues[row] == point && columnMinValues[column] == point {
				pairs = append(pairs, Pair{row + 1, column + 1})
			}
		}
	}

	return pairs
}
