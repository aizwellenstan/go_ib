package main

import (
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
	"flag"
	"strconv"
	"fmt"
)

func IntToStr(val int) string {
	return strconv.FormatFloat(float64(val), 'f', 2, 64)
}

func StrToInt(str string) int {
	toInt, _ := strconv.Atoi(str)
	return toInt
}

func StrToFloat(str string) float64 {
    f, err := strconv.ParseFloat(str, 64)
    if err != nil {
        fmt.Printf("%s\n", err.Error())
        return 0.0
    }
    return f
}

type DirtyTrades []struct {
	Op     string `json:"op"`
	Sl     string `json:"sl"`
	Symbol string `json:"symbol"`
	Time   string `json:"time"`
	Tp     string `json:"tp"`
	Vol    string `json:"vol"`
}

type Trade struct {
	Op     float64
	Sl     float64
	Symbol string
	Time   string
	Tp     float64
	Vol    int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Params defines accepted fields from command line
type Params struct {
	In        *string
	Out       *string
	Delimiter *string
	Fields    *int
}

func main() {
	file, err := ioutil.ReadFile("../../data/json/trades_2315.json")
	if err != nil {
		log.Println("ReadError: ", err)
		os.Exit(1)
	}

	var dirtyTrades DirtyTrades
	json.Unmarshal(file, &dirtyTrades)

	var interfaceSlice []interface{} = make([]interface{}, len(dirtyTrades))
	for i, _ := range dirtyTrades {
		trade := new(Trade)
		trade.Op = StrToFloat(dirtyTrades[i].Op)
		trade.Sl = StrToFloat(dirtyTrades[i].Sl)
		trade.Symbol = dirtyTrades[i].Symbol
		trade.Time = dirtyTrades[i].Time
		trade.Tp = StrToFloat(dirtyTrades[i].Tp)
		trade.Vol = StrToInt(dirtyTrades[i].Vol)
		interfaceSlice[i] = trade
	}

	params := Params{
		Out: flag.String("o", "../../data/json/trades_2315.json", "output file"),
	}

	finalJSON, marsherr := json.MarshalIndent(interfaceSlice, "", "\t")
	check(marsherr)

	ioerr := ioutil.WriteFile(*params.Out, finalJSON, 0644)
	check(ioerr)
}