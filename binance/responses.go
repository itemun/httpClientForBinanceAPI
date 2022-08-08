package binance

import (
	"fmt"
)

/*type assetsResponse struct {
	Assets []Asset `json:"data"`
	//Timestamp int64   `json:"timestamp"`
}*/

/*type assetResponse struct {
	Asset     Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}*/

type Asset struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int    `json:"openTime"`
	CloseTime          int    `json:"closeTime"`
	FirstId            int    `json:"firstId"`
	LastId             int    `json:"lastId"`
	Count              int    `json:"count"`
}

func (d Asset) Info() string {
	return fmt.Sprintf("[SYMBOL] %s | [LAST PRICE] %s | [PRICE CHANGE PERCENT] %s | [HIGH PRICE] %s | [LOW PRICE] %s | [VOLUME] %s", d.Symbol, d.LastPrice, d.PriceChangePercent, d.HighPrice, d.LowPrice, d.Volume)
}
