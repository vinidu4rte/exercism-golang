// Package weather is responsible to handle logic about weather.
package weather

// CurrentCondition stores current weather condition.
var CurrentCondition string

// CurrentLocation stores current city location.
var CurrentLocation string

// Forecast returns information about weather condition in city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
