package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
)

var sc = bufio.NewScanner(os.Stdin)
var wg = sync.WaitGroup{}
var rw = sync.RWMutex{}

func main() {
	fmt.Println("初期値を入力してください")
	xo := ReadFloat()
	fmt.Println("aの値を入力してください")
	a := ReadFloat()
	fmt.Println("繰り返し回数を入力してください")
	n := ReadInt()
	//result := logisticMapBig(big.NewFloat(xo), big.NewFloat(a), n)
	wg.Add(1)
	go func() {
		result := logisticMap(xo, a, n)
		judgeDemeanor(result[n-1], "ロジスティック写像")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		resultTent := tentMap(xo, a, n)
		judgeDemeanor(resultTent[n-1], "テント写像")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		resultBernouli := bernouliMap(xo, a, n)
		judgeDemeanor(resultBernouli[n-1], "ベルヌーイ写像")
		wg.Done()
	}()
	wg.Wait()
}

func judgeDemeanor(value float64, funcName string) {
	rw.Lock()
	fmt.Println(funcName + "は")
	if math.IsInf(value, -1) {
		fmt.Println("マイナス無限大に発散しました。\n")
	} else if math.IsInf(value, 1) {
		fmt.Println("無限大に発散しました。\n")
	} else {
		fmt.Printf("%vに収束しました。\n", value)
	}
	rw.Unlock()
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
