package main

import (
	"math/big"
)

func logisticMapBig(xo *big.Float, a *big.Float, n int) []big.Float {
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

func logisticMap(xo float64, a float64, n int) []float64 {
	array := make([]float64, n)
	array[0] = xo
	for i := 1; i < n; i++ {
		array[i] = a * array[i-1] * (1 - array[i-1])
	}
	return array
}

func tentMap(xo float64, a float64, n int) []float64 {
	array := make([]float64, n)
	array[0] = xo
	for i := 1; i < n; i++ {
		if array[i-1] < 0.5 {
			array[i] = a * array[i-1]
		} else if array[i-1] >= 0.5 {
			array[i] = a * (1 - array[i-1])
		} else {
			array[i] = 0
		}
	}
	return array
}

func bernouliMap(xo float64, a float64, n int) []float64 {
	array := make([]float64, n)
	array[0] = xo
	for i := 1; i < n; i++ {
		if array[i-1] < 0.5 {
			array[i] = a * array[i-1]
		} else if array[i-1] >= 0.5 && array[i-1] <= 1 {
			array[i] = a*array[i-1] - 1
		} else {
			array[i] = 0
		}
	}
	return array
}
