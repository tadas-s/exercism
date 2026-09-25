package microblog

func Truncate(phrase string) string {
	runes := []rune(phrase)
	return string(runes[0:min(5, len(runes))])
}
