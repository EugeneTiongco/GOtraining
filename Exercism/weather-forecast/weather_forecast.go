//Package weather provides the weather forecast of a location.
package weather

//CurrentCondition represents the weather condition of the city.
var CurrentCondition string

//CurrentLocation represents the location of the city.
var CurrentLocation string

//Forecast returns a string value detailing the weather condition of a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
