package main

import (
	"fmt"
	"github.com/pdepip/go-binance/binance"
	"os"
	"sync"
)

func callMe(curr string, wg sync.WaitGroup) {
	client := binance.New(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
	currency := binance.SymbolQuery{Symbol: curr}
	for {
		res, err := client.GetBookTicker(currency)
		if err != nil {
			panic(err)
		}
		fmt.Println(res.Symbol, "{bid: {price:", res.BidPrice, ", amount:", res.BidQuantity, "}, ask:{{price:", res.AskPrice, ", amount:", res.AskQuantity, "}}}")
	}
	defer wg.Done()
}

func main() {
	currencies := []string{
		"ETHUSDT",
		"BTCUSDT",
		"BNBUSDT",
	}
	var wg sync.WaitGroup
	wg.Add(len(currencies))
	for i := 0; i < len(currencies); i++ {
		go callMe(currencies[i], wg)
	}
	wg.Wait()
}
