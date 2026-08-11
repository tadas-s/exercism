package ledger

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

var locales = map[string]map[string]string{
	"en-US": {
		"date-format": "01/02/2006",
		"Date":        "Date",
		"Description": "Description",
		"Change":      "Change",
	},
	"nl-NL": {
		"date-format": "02-01-2006",
		"Date":        "Datum",
		"Description": "Omschrijving",
		"Change":      "Verandering",
	},
}

var currencies = map[string]string{
	"EUR": "€",
	"USD": "$",
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if _, localeExists := locales[locale]; !localeExists {
		return "", errors.New("invalid locale")
	}

	var entriesCopy []Entry

	for _, e := range entries {
		entriesCopy = append(entriesCopy, e)
	}

	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}

	slices.SortFunc(entriesCopy, func(a, b Entry) int {
		if n := cmp.Compare(a.Date, b.Date); n != 0 {
			return n
		}

		if n := cmp.Compare(a.Description, b.Description); n != 0 {
			return n
		}

		return cmp.Compare(a.Change, b.Change)
	})

	var s strings.Builder

	s.WriteString(fmt.Sprintf(
		"%-10s | %-25s | %-13s\n",
		locales[locale]["Date"],
		locales[locale]["Description"],
		locales[locale]["Change"],
	))

	for _, entry := range entriesCopy {
		parsedDate, err := time.Parse(time.DateOnly, entry.Date)

		if err != nil {
			return "", errors.New("bad date")
		}

		formattedDescription := entry.Description
		if len(formattedDescription) > 25 {
			formattedDescription = formattedDescription[:22] + "..."
		}

		formattedDate := parsedDate.Format(locales[locale]["date-format"])

		formattedCurrency, err := formatCurrency(locale, currency, entry.Change)

		if err != nil {
			return "", err
		}

		s.WriteString(
			fmt.Sprintf(
				"%-10s | %-25s | %13s\n",
				formattedDate,
				formattedDescription,
				formattedCurrency,
			),
		)
	}

	return s.String(), nil
}

func formatCurrency(locale string, currency string, cents int) (string, error) {
	var a string

	negative := cents < 0
	if negative {
		cents = -cents
	}

	if _, currencyExists := currencies[currency]; !currencyExists {
		return "", errors.New("bad currency")
	}

	currencySymbol := currencies[currency]

	if locale == "nl-NL" {
		a += currencySymbol
		a += " "
		centsStr := strconv.Itoa(cents)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}
		rest := centsStr[:len(centsStr)-2]
		var parts []string
		for len(rest) > 3 {
			parts = append(parts, rest[len(rest)-3:])
			rest = rest[:len(rest)-3]
		}
		if len(rest) > 0 {
			parts = append(parts, rest)
		}
		if negative {
			a += "-"
		}
		for i := len(parts) - 1; i >= 0; i-- {
			a += parts[i] + "."
		}
		a = a[:len(a)-1]
		a += ","
		a += centsStr[len(centsStr)-2:]
		a += " "
	} else if locale == "en-US" {
		if negative {
			a += "("
		}
		a += currencySymbol
		centsStr := strconv.Itoa(cents)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}
		rest := centsStr[:len(centsStr)-2]
		var parts []string
		for len(rest) > 3 {
			parts = append(parts, rest[len(rest)-3:])
			rest = rest[:len(rest)-3]
		}
		if len(rest) > 0 {
			parts = append(parts, rest)
		}
		for i := len(parts) - 1; i >= 0; i-- {
			a += parts[i] + ","
		}
		a = a[:len(a)-1]
		a += "."
		a += centsStr[len(centsStr)-2:]
		if negative {
			a += ")"
		} else {
			a += " "
		}
	} else {
		return "", errors.New("bad locale")
	}

	return a, nil
}
