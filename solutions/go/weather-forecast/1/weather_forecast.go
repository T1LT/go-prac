// Package weather provides functions to
// get the weather forecast.
package weather

// CurrentCondition is a string that stores the present weather condition.
var CurrentCondition string
// CurrentLocation is a string that stores the location of the place that needs a weather forecast.
var CurrentLocation string

// Forecast takes in a city and a condition and returns the forecast based on the arguments.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
