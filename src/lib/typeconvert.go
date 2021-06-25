package lib

import (
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