package romannumerals

import (
	"fmt"
	"math"
	"strings"
)

var units = [][]string{
	{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
	{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
	{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	{"M", "MM", "MMM"},
}

func ToRomanNumeral(input int) (string, error) {
	var roman strings.Builder

	if input < 1 || input > 3999 {
		return "", fmt.Errorf("only numbers from 1 to 3999 supported, got %d", input)
	}

	for i := 3; i >= 0; i-- {
		digit := (input % int(math.Pow(10, float64(i+1)))) / int(math.Pow(10, float64(i)))
		if digit > 0 {
			roman.WriteString(units[i][digit-1])
		}
	}

	return roman.String(), nil
}
