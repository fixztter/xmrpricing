package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	api    = "https://api.nanopool.org/v1/xmr/prices"
	client *http.Client
)

type XMRPrice struct {
	PriceBTC float32 `json:"price_btc"`
	PriceUSD float32 `json:"price_usd"`
	PriceRUR float32 `json:"price_rur"`
	PriceEUR float32 `json:"price_eur"`
	PriceCNY float32 `json:"price_cny"`
	PriceGBP float32 `json:"price_gbp"`
}

type XMRCurrency struct {
	Status bool     `json:"status"`
	Data   XMRPrice `json:"data"`
}

func (x *XMRCurrency) GetMoneroPrice(url string) (XMRCurrency, error) {
	var xmrPrice XMRCurrency

	if err := getJson(url, &xmrPrice); err != nil {
		return xmrPrice, err
	}

	return xmrPrice, nil
}

func (x XMRPrice) String() string {
	return fmt.Sprintf("Price BTC: %.8f\nPrice USD: %.2f\nPrice RUR: %.2f\nPrice EUR: %.2f\nPrice CNY: %.2f\nPriceGBP: %.2f\n",
		x.PriceBTC, x.PriceUSD, x.PriceRUR, x.PriceEUR, x.PriceCNY, x.PriceGBP)
}

func (x XMRCurrency) String() string {
	return fmt.Sprintf("Status: %v\nData: %s\n", x.Status, x.Data)
}

func main() {
	client = &http.Client{Timeout: time.Second * 10}
	var xmrPrice XMRCurrency
	data, err := xmrPrice.GetMoneroPrice(api)

	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err.Error())
		return
	}

	fmt.Printf("XMR: %.2f\n", data.Data.PriceUSD)
}

func getJson(url string, target any) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&target)
}
