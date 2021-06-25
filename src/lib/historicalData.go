package lib

import (
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
)

type HisBars []struct {
	S string `json:"s"`
	BarData []BarData `json:"d"`
}

type BarData struct {
	Date string  `json:"d"`
	Open float64 `json:"o"`
	High float64 `json:"h"`
	Low float64 `json:"l"`
	Close float64 `json:"c"`
}


func LoadBarData() HisBars {
	file, err := ioutil.ReadFile("./data/json/hisBarsStocksD1arr.json")
	if err != nil {
		log.Println("ReadError: ", err)
		os.Exit(1)
	}

	var hisbars HisBars
	json.Unmarshal(file, &hisbars)

	return hisbars
}