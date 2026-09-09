package spaceage

type Planet string

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Mercury":
		return seconds / 7600543.8
	case "Venus":
		return seconds / 19414149.1
	case "Earth":
		return seconds / 31557600.0
	case "Mars":
		return seconds / 59354032.7
	case "Jupiter":
		return seconds / 374355659.1
	case "Saturn":
		return seconds / 929292362.9
	case "Uranus":
		return seconds / 2651370019.3
	case "Neptune":
		return seconds / 5200418560.0
	default:
		return -1
	}
}
