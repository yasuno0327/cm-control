package main

import (
	"bufio"
	"fmt"
	"math"
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
	//result := logisticMapBig(big.NewFloat(xo), big.NewFloat(a), n)
	result := logisticMap(xo, a, n)
	if math.IsInf(result[n-1], -1) {
		fmt.Println("マイナス無限大に発散しました。")
	} else if math.IsInf(result[n-1], 1) {
		fmt.Println("無限大に発散しました。")
	} else {
		fmt.Printf("%vに収束しました。", result[n-1])
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
