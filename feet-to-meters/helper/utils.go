package helper

// Feet2Meter converts from feet to meters
func Feet2Meter(feet float64) float64 {
	meter := feet * 0.3048
	return meter
}

// Celsius2Farenheit converts from celsius to fareinheit
func Celsius2Farenheit(celsius float64) float64 {
	farenheit := (celsius * 9. / 5) + 32.
	return farenheit
}
