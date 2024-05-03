# Get Monero Price

Get the price of Monero with this little tool.

## Supports Multiple Currencies

The information is obtained from [here](https://api.nanopool.org/v1/xmr/prices).

Supports the next currencies:

- Bitcoin (BTC)
- United States Dollar (USD)
- Euro (EUR)
- Russian Ruble (RUR)
- Chinese Yuan (CNY)
- British Pound Sterling (GBP)

## Build from source

You can build with a specific name:

```sh
    go build -o <name> ./main.go
```

## Modify the code at your convenience

Get the output you want:

```go
    // From this
    fmt.Printf("XMR: %.2f\n", data.Data.PriceUSD)
    // To this
    fmt.Printf("XMR: %.2f\n", data.Data.PriceEUR)
  
```
##
