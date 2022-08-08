package main

import (
	"fmt"
	"httpClient/binance"
	"log"
	"time"
)

func main() {
	binanceClient, err := binance.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	/*	assets, err := binanceClient.GetAssets()
		if err != nil {
			log.Fatal(err)
		}
		for _, asset := range assets {
			fmt.Println(asset.Info())
		}*/

	bitcoin, err := binanceClient.GetAsset("BTCUSDT")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bitcoin.Info())
}
