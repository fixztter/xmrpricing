package xmrpricing

import (
	"fmt"

	"github.com/fixztter/xmrpricing/internal/json"
)

// api: API that provide's XMR pricing
const api = "https://api.nanopool.org/v1/xmr/prices"

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

func (x *XMRCurrency) GetMoneroPrice() (*XMRCurrency, error) {
	if err := json.GetJSON(api, x); err != nil {
		return x, err
	}

	return x, nil
}

func (x XMRPrice) String() string {
	return fmt.Sprintf("Price BTC: %.8f\nPrice USD: %.2f\nPrice RUR: %.2f\nPrice EUR: %.2f\nPrice CNY: %.2f\nPriceGBP: %.2f\n",
		x.PriceBTC, x.PriceUSD, x.PriceRUR, x.PriceEUR, x.PriceCNY, x.PriceGBP)
}

func (x XMRCurrency) String() string {
	return fmt.Sprintf("Status: %v\nData: %s\n", x.Status, x.Data)
}
