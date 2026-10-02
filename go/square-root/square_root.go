package squareroot

func SquareRoot(number int) (int, error) {
	res := 0
	bit := 1 << 62

	for bit > number {
		bit >>= 2
	}

	for bit != 0 {
		if number >= res+bit {
			number -= res + bit
			res = (res >> 1) + bit
		} else {
			res >>= 1
		}

		bit >>= 2
	}

	return res, nil
}
