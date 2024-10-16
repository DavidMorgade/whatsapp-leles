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

	"github.com/joho/godotenv"
)

type Quote struct {
	EUR struct {
		Price                 float64 `json:"price"`
		Volume24h             float64 `json:"volume_24h"`
		VolumeChange24h       float64 `json:"volume_change_24h"`
		PercentChange1h       float64 `json:"percent_change_1h"`
		PercentChange24h      float64 `json:"percent_change_24h"`
		PercentChange7d       float64 `json:"percent_change_7d"`
		PercentChange30d      float64 `json:"percent_change_30d"`
		MarketCap             float64 `json:"market_cap"`
		MarketCapDominance    float64 `json:"market_cap_dominance"`
		FullyDilutedMarketCap float64 `json:"fully_diluted_market_cap"`
		LastUpdated           string  `json:"last_updated"`
	} `json:"EUR"`
}

type CryptoData struct {
	ID                            int         `json:"id"`
	Name                          string      `json:"name"`
	Symbol                        string      `json:"symbol"`
	Slug                          string      `json:"slug"`
	IsActive                      int         `json:"is_active"`
	IsFiat                        int         `json:"is_fiat"`
	CirculatingSupply             float64     `json:"circulating_supply"`
	TotalSupply                   float64     `json:"total_supply"`
	MaxSupply                     float64     `json:"max_supply"`
	DateAdded                     string      `json:"date_added"`
	NumMarketPairs                int         `json:"num_market_pairs"`
	CMCRank                       int         `json:"cmc_rank"`
	LastUpdated                   string      `json:"last_updated"`
	Tags                          []string    `json:"tags"`
	Platform                      interface{} `json:"platform"`
	SelfReportedCirculatingSupply interface{} `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         interface{} `json:"self_reported_market_cap"`
	Quote                         Quote       `json:"quote"`
}

type Response struct {
	Data   map[string]CryptoData `json:"data"`
	Status struct {
		Timestamp    string `json:"timestamp"`
		ErrorCode    int    `json:"error_code"`
		ErrorMessage string `json:"error_message"`
		Elapsed      int    `json:"elapsed"`
		CreditCount  int    `json:"credit_count"`
		Notice       string `json:"notice"`
	} `json:"status"`
}

func init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GetCryptoPrice(crypto string) (string, error) {
	API_KEY := os.Getenv("COIN_API_KEY")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
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

	var response Response
	err = json.Unmarshal(respBody, &response)

	if err != nil {
		return "", err
	}

	// Access the first (and only) element in the data field
	for _, coin := range response.Data {
		result := fmt.Sprintf("Symbol: %s, Name: %s, Price: %.8f EUR", coin.Symbol, coin.Name, coin.Quote.EUR.Price)
		fmt.Println(result)
		return result, nil
	}

	return "Crypto not found", nil
}
