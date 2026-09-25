package reversestring

func Reverse(input string) string {
	out := []rune(input)

	for i := 0; i < len(out)/2; i++ {
		out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
	}

	return string(out)
}
