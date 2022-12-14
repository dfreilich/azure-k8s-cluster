package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	url          = "https://blockchain.info/ticker"
	printPrice   = time.Minute * 2
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

func printBitcoinInLoop() error {
	vals := []float64{}
	sum := 0.0

	idx := 0
	for true {
		idx++
		val := getBitcoinEuro()
		fmt.Printf("Current Value: %v\n", val)

		if idx == printAverage {
			average := (sum / float64(len(vals)))
			fmt.Printf("Average: %v\n", average)

			var first float64
			first, vals = vals[0], vals[1:]
			sum -= first
			idx = 0
		}

		vals = append(vals, val)
		sum += val

		time.Sleep(printPrice)
	}
	return nil
}

func getBitcoinEuro() float64 {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	bitcoin := &Bitcoin{}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, bitcoin)
	if err != nil {
		fmt.Println(err)
	}

	return bitcoin.Eur.Last
}
