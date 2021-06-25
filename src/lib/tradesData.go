package lib

import (
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
)

type Trades []struct {
	Op     float64
	Sl     float64
	Symbol string
	Time   string
	Tp     float64
	Vol    int
}


func LoadTradeData() Trades {
	file, err := ioutil.ReadFile("./data/json/trades_2315.json")
	if err != nil {
		log.Println("ReadError: ", err)
		os.Exit(1)
	}

	var trades Trades
	json.Unmarshal(file, &trades)

	return trades
}