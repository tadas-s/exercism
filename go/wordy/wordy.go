package wordy

import (
	"fmt"
	"regexp"
	"strconv"
)

type TokenType int

const (
	TokenError TokenType = iota
	TokenOperand
	TokenOperatorPlus
	TokenOperatorMinus
	TokenOperatorMultiply
	TokenOperatorDivide
	TokenFluff
)

type Token struct {
	Type    TokenType
	Operand int
}

func Answer(question string) (int, bool) {
	var parsed []Token

	r, _ := regexp.Compile("(What is|\\?|plus|minus|multiplied by|divided by|-?\\d+|.+)\\s*")
	matches := r.FindAllStringSubmatch(question, -1)

	for _, match := range matches {
		parsedToken, err := parseToken(match[1])

		if err != nil {
			return 0, false
		}

		if parsedToken.Type != TokenFluff {
			parsed = append(parsed, parsedToken)
		}
	}

	answer, err := calculate(parsed)

	if err != nil {
		return 0, false
	}

	return answer, true
}

func parseToken(token string) (Token, error) {
	if ok, _ := regexp.MatchString("^-?\\d+$", token); ok {
		if operand, err := strconv.Atoi(token); err == nil {
			return Token{TokenOperand, operand}, nil
		}

		return Token{TokenError, 0}, fmt.Errorf("error parsing numeric operand '%s'", token)
	}

	switch token {
	case "What is", "?":
		return Token{TokenFluff, 0}, nil
	case "plus":
		return Token{TokenOperatorPlus, 0}, nil
	case "minus":
		return Token{TokenOperatorMinus, 0}, nil
	case "multiplied by":
		return Token{TokenOperatorMultiply, 0}, nil
	case "divided by":
		return Token{TokenOperatorDivide, 0}, nil
	}

	return Token{TokenError, 0}, fmt.Errorf("error parsing token '%s", token)
}

func calculate(tokens []Token) (int, error) {
	stackIndex := 0

	if len(tokens) < 1 {
		return 0, fmt.Errorf("nothing to do")
	}

	for {
		t1 := tokens[stackIndex]
		stackIndex++

		if t1.Type == TokenOperand && stackIndex == len(tokens) {
			return t1.Operand, nil
		}

		if stackIndex == len(tokens) {
			return 0, fmt.Errorf("unexpected end of operations list")
		}

		t2 := tokens[stackIndex]
		stackIndex++

		if t1.Type == TokenOperand && t2.Type >= TokenOperatorPlus && t2.Type <= TokenOperatorDivide {
			if stackIndex == len(tokens) {
				return 0, fmt.Errorf("unexpected end of operations list")
			}

			t3 := tokens[stackIndex]
			stackIndex++

			if t3.Type != TokenOperand {
				return 0, fmt.Errorf("unexpected token")
			}

			opResult := 0

			switch t2.Type {
			case TokenOperatorPlus:
				opResult = t1.Operand + t3.Operand
			case TokenOperatorMinus:
				opResult = t1.Operand - t3.Operand
			case TokenOperatorMultiply:
				opResult = t1.Operand * t3.Operand
			case TokenOperatorDivide:
				opResult = t1.Operand / t3.Operand
			default:
				break
			}

			stackIndex--
			tokens[stackIndex] = Token{TokenOperand, opResult}
		} else {
			return 0, fmt.Errorf("unexpected token")
		}
	}
}
