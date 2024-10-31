package main

import (
	"fmt"

	"github.com/fixztter/xmrpricing"
)

func main() {
	var xmrPrice xmrpricing.XMRCurrency
	data, err := xmrPrice.GetMoneroPrice()

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	fmt.Printf("XMR: %.2f\n", data.Data.PriceUSD)
}
