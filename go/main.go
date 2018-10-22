package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("初期値を入力してください")
	xo := ReadFloat()
	fmt.Println("aの値を入力してください")
	a := ReadFloat()
	fmt.Println("繰り返し回数を入力してください")
	n := ReadInt()
	result := logisticMap(big.NewFloat(xo), big.NewFloat(a), n)
	for i := range result {
		fmt.Println(result[i].String())
	}
}

func ReadFloat() float64 {
	sc.Scan()
	v, _ := strconv.ParseFloat(sc.Text(), 64)
	return v
}

func ReadInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
