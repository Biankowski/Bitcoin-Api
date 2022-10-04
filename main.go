package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Bitcoin struct {
	Base     string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}
type Data struct {
	InnerData Bitcoin `json:"data"`
}

func main() {
	url := "https://api.coinbase.com/v2/prices/spot?currency=USD"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var bitCoinResponse Data
	json.Unmarshal([]byte(data), &bitCoinResponse)
	fmt.Printf("Base: %s\nCurrency: %s\nAmount: %s\n", bitCoinResponse.InnerData.Base, bitCoinResponse.InnerData.Currency, bitCoinResponse.InnerData.Amount)

}
