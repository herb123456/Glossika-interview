package utils

import (
	"math"
	"math/rand"
	"time"
)

func RandomDigits(num float64) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(int(math.Pow(10.0, num)))
}
