package isbnverifier

func IsValidISBN(isbn string) bool {
	var digitsCollected = 0
	var sum = 0
	var lastPos = len(isbn) - 1

	for i, char := range isbn {
		if i <= lastPos && char >= '0' && char <= '9' {
			sum += int(char-'0') * (10 - digitsCollected)
			digitsCollected++
		} else if i == lastPos && char == 'X' {
			sum += 10
			digitsCollected++
		} else if char == '-' && (digitsCollected == 1 || digitsCollected == 4 || digitsCollected == 9) {
			continue
		} else {
			return false
		}
	}

	return digitsCollected == 10 && sum%11 == 0
}
