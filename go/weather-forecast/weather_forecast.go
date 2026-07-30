// Package weather provides utilities for formatting weather forecasts.
package weather

var (
	// CurrentCondition keeps track of most recently used weather condition.
	CurrentCondition string
	// CurrentLocation keeps track of most recently used location.
	CurrentLocation string
)

// Forecast formats a weather forecast string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
