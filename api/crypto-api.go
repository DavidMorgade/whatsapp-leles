package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Response struct {
	Data []struct {
		Name  string `json:"name"`
		Quote struct {
			USD struct {
				Price float64 `json:"price"`
			} `json:"USD"`
		} `json:"quote"`
	} `json:"data"`
}

func GetCryptoPrice(crypto string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	respBody, _ := io.ReadAll(resp.Body)

	var response Response
	err = json.Unmarshal([]byte(respBody), &response)

	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	for _, coin := range response.Data {
		if coin.Name == strings.ReplaceAll(crypto, " ", "") {
			fmt.Printf("Name: %s, Price: %.8f USD\n", coin.Name, coin.Quote.USD.Price)
			return
		}
	}

	fmt.Printf("Cryptocurrency %s not found\n", crypto)
}
