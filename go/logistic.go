package main

import (
	"math/big"
)

func logisticMap(xo *big.Float, a *big.Float, n int) []big.Float {
	array := make([]big.Float, n)
	array[0] = *xo
	for i := 1; i < n; i++ {
		z := new(big.Float)
		result := new(big.Float)
		z.Mul(a, &array[i-1])
		result.Mul(z, new(big.Float).Sub(big.NewFloat(1), &array[i-1]))
		array[i] = *result
	}
	return array
}
