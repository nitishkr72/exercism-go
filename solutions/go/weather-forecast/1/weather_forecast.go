// Package weather exports Forecast function that will give you information
// about current location and weather condition.
package weather

var (
    // CurrentCondition stores the current weather condition.
	CurrentCondition string
    // CurrentLocation stores current locaation.
	CurrentLocation  string
)

// Forecast function takes city name, condition as string and return
// current weather forecast at that station.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
