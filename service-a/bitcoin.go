package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

var (
	url          = "https://blockchain.info/ticker"
	checkPrice   = time.Minute * 1
	printAverage = 10
)

type Bitcoin struct {
	Eur struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"EUR"`
}

func getBitcoinEuro() (float64, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, errors.Wrapf(err, "getting url %s", url)
	}
	bitcoin := &Bitcoin{}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, bitcoin)
	if err != nil {
		return 0, errors.Wrapf(err, "unmarshelling response from url %s", url)
	}

	return bitcoin.Eur.Last, nil
}

func printBitcoinAndAverage() error {
	ticker := time.NewTicker(checkPrice)
	defer ticker.Stop()

	var prices []float64
	for {
		select {
		case <-ticker.C:
			price, err := getBitcoinEuro()
			if err != nil {
				fmt.Println(err)
				continue
			}

			prices = append(prices, price)
			fmt.Printf("Current Bitcoin price in Euro: %.2f\n", price)

			if len(prices) == printAverage {
				fmt.Printf("Average Bitcoin price in Euro over the last 10 minutes: %.2f\n", avg(prices))
				prices = []float64{}
			}
		}
	}
	return nil
}

func avg(vals []float64) float64 {
	sum := 0.0
	for _, val := range vals {
		sum += val
	}
	return sum / float64(len(vals))
}
