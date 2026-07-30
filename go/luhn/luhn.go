package luhn

func Valid(id string) bool {
	var sum int64 = 0
	var digits uint = 0

	for i := len(id) - 1; i >= 0; i-- {
		if id[i] >= '0' && id[i] <= '9' {
			var digit = id[i] - '0'

			if digits%2 > 0 {
				if digit*2 > 9 {
					sum += int64(digit*2 - 9)
				} else {
					sum += int64(digit * 2)
				}
			} else {
				sum += int64(digit)
			}

			digits++
		} else if id[i] == ' ' {
			continue // skip spaces
		} else {
			return false
		}
	}

	return digits > 1 && sum%10 == 0
}
