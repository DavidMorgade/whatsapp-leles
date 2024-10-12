package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/whatsapp-leles/utils"
)

var apiURL = "https://api.openweathermap.org/data/2.5/weather?lang=es&lat=15.0286&lon=120.6898&appid="

type Weather struct {
	Cod  CodType `json:"cod"`
	Name string  `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
}
type CodType struct {
	Value int
}

func (c *CodType) UnmarshalJSON(data []byte) error {
	var intValue int
	if err := json.Unmarshal(data, &intValue); err == nil {
		c.Value = intValue
		return nil
	}

	var stringValue string
	if err := json.Unmarshal(data, &stringValue); err == nil {
		var err error
		c.Value, err = strconv.Atoi(stringValue)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("invalid cod value")
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	apiURL += apiKey
}

func GetWeather() (*Weather, error) {
	resp, err := http.Get(apiURL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if weather.Cod.Value == 404 {
		return nil, errors.New("Ciudad no encontrada")
	}
	if err != nil {
		return nil, err
	}

	fmt.Printf("City: %s\n", weather.Name)
	fmt.Printf("Temperature: %.2f\n", utils.KelvinToCelsius(weather.Main.Temp))
	fmt.Printf("Weather Description: %s\n", weather.Weather[0].Description)
	fmt.Printf("Wind Speed: %.2f\n", weather.Wind.Speed)
	fmt.Printf("Cloud Percentage: %d\n", weather.Clouds.All)

	return &weather, nil
}

func GetWeatherByCity(city string) (*Weather, error) {
	if len(strings.ReplaceAll(city, " ", "")) == 0 {
		city = "san fernando"
	}
	apiURL := "https://api.openweathermap.org/data/2.5/weather?lang=es&q=" + city + "&appid="
	apiKey := os.Getenv("API_KEY")
	apiURL += apiKey
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if weather.Cod.Value == 404 {
		return nil, errors.New("Ciudad no encontrada")
	}
	if err != nil {
		return nil, err
	}

	fmt.Printf("City: %s\n", weather.Name)
	fmt.Printf("Temperature: %.2f\n", utils.KelvinToCelsius(weather.Main.Temp))
	fmt.Printf("Weather Description: %s\n", weather.Weather[0].Description)
	fmt.Printf("Wind Speed: %.2f\n", weather.Wind.Speed)
	fmt.Printf("Cloud Percentage: %d\n", weather.Clouds.All)

	return &weather, nil
}
