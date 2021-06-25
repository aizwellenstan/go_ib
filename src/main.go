package main

import (
	"fmt"
	."github.com/aizwellenstan/go_ib/src/lib"
)


func main(){
	arr := [] float64{1,2}
	sma25 := Sma(arr,2)
	fmt.Println(sma25[1])

	tp := 1.007
	tp = NormalizeFloat(tp,1.02,1.03)
	fmt.Println(tp)

	hisBarsD1 := LoadBarData()
	fmt.Println(hisBarsD1[0].BarData[0].Date)

	slope := Slope(hisBarsD1[0].BarData, &SlopeOpts{})
	fmt.Println(slope)

	trades := LoadTradeData()
	fmt.Println(trades)
}