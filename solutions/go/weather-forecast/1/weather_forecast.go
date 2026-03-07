// Package weather provides tools to get weather information.
package weather

var (
	// CurrentCondition declares the current weather condition.
	CurrentCondition string
	// CurrentLocation declares the current location.
	CurrentLocation string
)

// Forecast takes the current city and weather condition, and returns
// a formatted forecast string value.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
