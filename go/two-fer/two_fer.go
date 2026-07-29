// Package twofer generates blurbs to say when sharing two-for-one purchases.
package twofer

import "fmt"

// ShareWith builds a blurb to say when sharing a two-for-one purchase.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
