package allergies

var allergens = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(allergies uint) []string {
	var result []string

	for allergen := range allergens {
		if AllergicTo(allergies, allergen) {
			result = append(result, allergen)
		}
	}

	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	if bitmask, ok := allergens[allergen]; ok {
		return (allergies & bitmask) != 0
	}

	return false
}
