package binarysearch

func SearchInts(list []int, key int) int {
	lowerBound := 0
	upperBound := len(list) - 1

	for lowerBound <= upperBound {
		i := (lowerBound + upperBound) / 2

		if list[i] < key {
			lowerBound = i + 1
		} else if list[i] > key {
			upperBound = i - 1
		} else {
			return i
		}
	}

	return -1
}
