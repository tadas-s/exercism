package ledger

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
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
	negative := cents < 0
	if negative {
		cents = -cents
	}

	if _, currencyExists := currencies[currency]; !currencyExists {
		return "", errors.New("bad currency")
	}

	switch {
	case locale == "nl-NL" && negative:
		return fmt.Sprintf("%s -%s ", currencies[currency], formatCents(cents, ".", ",")), nil
	case locale == "nl-NL":
		return fmt.Sprintf("%s %s ", currencies[currency], formatCents(cents, ".", ",")), nil
	case locale == "en-US" && negative:
		return fmt.Sprintf("(%s%s)", currencies[currency], formatCents(cents, ",", ".")), nil
	case locale == "en-US":
		return fmt.Sprintf("%s%s ", currencies[currency], formatCents(cents, ",", ".")), nil
	default:
		return "", errors.New("bad locale")
	}
}

func formatCents(cents int, thousandsSeparator, decimalSeparator string) string {
	centsStr := fmt.Sprintf("%03d", cents)
	rest := centsStr[:len(centsStr)-2]

	var parts []string
	for len(rest) > 3 {
		parts = append(parts, rest[len(rest)-3:])
		rest = rest[:len(rest)-3]
	}
	if len(rest) > 0 {
		parts = append(parts, rest)
	}
	slices.Reverse(parts)

	return strings.Join(parts, thousandsSeparator) + decimalSeparator + centsStr[len(centsStr)-2:]
}
