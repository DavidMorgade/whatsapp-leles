package utils

func KelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func GetCityFromMessage(message string) string {
	return message[9:]
}
