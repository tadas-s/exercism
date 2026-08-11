package phonenumber

import (
	"errors"
	"fmt"
)

func Number(phoneNumber string) (string, error) {
	cleanedNumber := ""

	for _, digit := range phoneNumber {
		if digit >= '0' && digit <= '9' {
			cleanedNumber += string(digit)
		}
	}

	if len(cleanedNumber) == 11 && cleanedNumber[0] == '1' {
		cleanedNumber = cleanedNumber[1:]
	}

	if len(cleanedNumber) >= 11 {
		return "", errors.New("phone number too long")
	}

	if len(cleanedNumber) < 10 {
		return "", errors.New("phone number is less than 9 digits")
	}

	if cleanedNumber[0] == '0' || cleanedNumber[0] == '1' {
		return "", errors.New("invalid area code")
	}

	if cleanedNumber[3] == '0' || cleanedNumber[3] == '1' {
		return "", errors.New("invalid exchange code")
	}

	return cleanedNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	cleanedNumber, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}

	return cleanedNumber[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	cleanedNumber, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
			"(%s) %s-%s",
			cleanedNumber[0:3],
			cleanedNumber[3:6],
			cleanedNumber[6:]),
		nil
}
