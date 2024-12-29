package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fixztter/xmrpricing"
)

func main() {
	// defaults
	//var xmrPrice xmrpricing.XMRCurrency
	//data, err := xmrPrice.GetMoneroPrice()

	//if err != nil {
	//	fmt.Printf("error: %s\n", err.Error())
	//	return
	//}

	//fmt.Printf("XMR: %.2f\n", data.Data.PriceUSD)

	//xmrPrice := new(xmrpricing.XMRCurrency)

	flag.Usage = func() {
		f := flag.CommandLine.Output()
		fmt.Fprintln(f, "xmrpricing - Get monero (XMR) pricing")
		fmt.Fprintln(f, "")
		fmt.Fprintf(f, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintln(f, "Options:")
		flag.PrintDefaults()
	}

	currency := flag.String(
		"currency",
		"",
		fmt.Sprintf("%s [%s]\n%s\n",
			"get monero (XMR) price in a specific currency",
			"btc|usd|eur|rur|cny|gbp",
			"example: xmrpicing -currency=btc",
		),
	)

	flag.Parse()

    dt, err := getCurrencyData()
    if err != nil {
        log.Fatalf("error: %s\n", err.Error())
    }

	switch *currency {
	case "btc":
        fmt.Printf("BTC: %.6f\n", dt.Data.PriceBTC)
	case "usd":
        fmt.Printf("USD: %.2f\n", dt.Data.PriceUSD)
	case "eur":
        fmt.Printf("EUR: %.2f\n", dt.Data.PriceEUR)
	case "rur":
        fmt.Printf("RUR: %.2f\n", dt.Data.PriceRUR)
	case "cny":
        fmt.Printf("CNY: %.2f\n", dt.Data.PriceCNY)
	case "gbp":
        fmt.Printf("GBP: %.2f\n", dt.Data.PriceGBP)
	default:
		fmt.Printf("you need to give a specific currency\n")
        flag.PrintDefaults()
	}
}

func getCurrencyData() (*xmrpricing.XMRCurrency, error) {
	xmrPrice := new(xmrpricing.XMRCurrency)
	dt, err := xmrPrice.GetMoneroPrice()
	if err != nil {
		return nil, err
	}

	return dt, err
}
