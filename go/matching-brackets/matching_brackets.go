package matchingbrackets

import "errors"

type runeStack []rune

func (s *runeStack) Push(r rune) {
	*s = append(*s, r)
}

func (s *runeStack) Pop() (rune, error) {
	if len(*s) < 1 {
		return 0, errors.New("empty stack")
	}

	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return r, nil
}

func openingBracket(r rune) (rune, error) {
	switch r {
	case ']':
		return '[', nil
	case '}':
		return '{', nil
	case ')':
		return '(', nil
	default:
		return 0, errors.New("bad bracket rune")
	}
}

func Bracket(input string) bool {
	var stack runeStack

	for _, r := range input {
		switch r {
		case '[', '{', '(':
			stack.Push(r)
		case ']', '}', ')':
			expectedBracket, _ := openingBracket(r)

			if actual, err := stack.Pop(); err == nil {
				if actual != expectedBracket {
					return false
				}
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
