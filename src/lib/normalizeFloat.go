package lib

import (
	"math"
	"strings"
	"strconv"
)

func NormalizeFloat(price, sample1, sample2 float64) float64 {
	dec1 := numDecPlaces(sample1)
	dec2 := numDecPlaces(sample2)
	dec := max(dec1,dec2)
	res := toFixed(price,dec)
	return res
}

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}

func numDecPlaces(v float64) int {
    s := strconv.FormatFloat(v, 'f', -1, 64)
    i := strings.IndexByte(s, '.')
    if i > -1 {
        return len(s) - i - 1
    }
    return 0
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
    