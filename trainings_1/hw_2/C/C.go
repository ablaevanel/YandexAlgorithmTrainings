package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var num float64
	fmt.Scan(&num)
	minAbs := math.Abs(num - arr[0])
	ans := 0
	for i := 1; i < n; i++ {
		if math.Abs(num-arr[i]) < minAbs {
			ans = i
			minAbs = math.Abs(num - arr[i])
		}
	}
	fmt.Print(arr[ans])
}
