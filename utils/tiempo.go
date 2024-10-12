package utils

import (
	"fmt"

	"github.com/whatsapp-leles/api"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func meterSecondToKilometerHour(ms float64) float64 {
	return ms * 3.6
}

func GetCityFromMessage(message string) string {
	return message[9:]
}

func SendWeatherMessage(weather api.Weather, client *whatsmeow.Client, v *events.Message) {
	country := countryCodeToFullNameOnSpanish(weather.Sys.Country)

	SendMessage("Ciudad: "+weather.Name, client, v)
	SendMessage("País: "+country, client, v)
	SendMessage("Temperatura: "+fmt.Sprintf("%.2f", kelvinToCelsius(weather.Main.Temp)), client, v)
	SendMessage("Descripción del clima: "+weather.Weather[0].Description, client, v)
	SendMessage("Velocidad del viento: "+fmt.Sprintf("%.2f", meterSecondToKilometerHour(weather.Wind.Speed)), client, v)
}
