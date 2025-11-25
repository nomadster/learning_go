//Package weather fa cose belle.
package weather

var (
    //CurrentCondition represents the current weather condition.
	CurrentCondition string
    //CurrentLocation represents the name of the current location.
	CurrentLocation  string
)

//Forecast forecasts weather condition for given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
