package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Response struct {
	Data map[string]interface{} `json:"data"`
}

func init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GetCryptoPrice(crypto string) (string, error) {

	crypto = strings.ReplaceAll(crypto, " ", "")

	if crypto == "" {
		return "", errors.New("No se ha especificado ninguna criptomoneda")
	}

	API_KEY := os.Getenv("COIN_API_KEY")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest", nil)
	if err != nil {
		return "", err
	}

	q := url.Values{}
	q.Add("symbol", strings.ReplaceAll(strings.ToUpper(crypto), " ", ""))
	if len(q.Get("symbol")) > 5 {
		q.Del("symbol")
		q.Add("slug", strings.ReplaceAll(strings.ToLower(crypto), " ", ""))
	}
	q.Add("convert", "EUR")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", API_KEY)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("No se encontró la criptomoneda %s", crypto)
	}

	var response Response
	err = json.Unmarshal([]byte(respBody), &response)

	if err != nil {
		return "", err
	}

	for _, data := range response.Data {
		switch v := data.(type) {
		case map[string]interface{}:
			lastUpdated, err := convertDate(v["last_updated"].(string))
			if err != nil {
				return "", err
			}
			name := v["name"].(string)
			symbol := v["symbol"].(string)
			quote := v["quote"].(map[string]interface{})
			eur := quote["EUR"].(map[string]interface{})
			price := eur["price"].(float64)
			return fmt.Sprintf("El precio de %s %s es de %.8f€, ultima actualizacion a %s", name, symbol, price, lastUpdated), nil
		case []interface{}:
			if len(v) > 0 {
				lastUpdated, err := convertDate(v[0].(map[string]interface{})["last_updated"].(string))
				if err != nil {
					return "", err
				}
				name := v[0].(map[string]interface{})["name"].(string)
				symbol := v[0].(map[string]interface{})["symbol"].(string)
				coin := v[0].(map[string]interface{})
				quote := coin["quote"].(map[string]interface{})
				eur := quote["EUR"].(map[string]interface{})
				price := eur["price"].(float64)
				return fmt.Sprintf("El precio de %s %s es de %.8f€, ultima actualizacion a %s", name, symbol, price, lastUpdated), nil
			}
		}
	}

	return "No se encontró la criptomoneda", nil
}

func convertDate(input string) (string, error) {
	// Parse the input date string
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return "", err
	}

	// Format the date into the desired format
	output := t.Format("02/01/2006 15:04:05")
	return output, nil
}
