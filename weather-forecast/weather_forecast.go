// Package weather provides the current weather condition of various cities in Globinocus.
package weather

// CurrentConditions represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current location or city of Globinocus.
var CurrentLocation string

// Forecast returns the current weather condition of the city informed in CurrentLocation variable.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
