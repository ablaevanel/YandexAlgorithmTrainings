package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var ansLen int
	var ans []int
	for i := 0; i < n; i++ {
		start, end := i, n-1
		for start < end && arr[start] == arr[end] {
			start++
			end--
		}
		if start >= end {
			ansLen = i
			for k := i - 1; k >= 0; k-- {
				ans = append(ans, arr[k])
			}
			break
		}
	}
	fmt.Println(ansLen)
	for i := 0; i < ansLen; i++ {
		fmt.Print(ans[i], " ")
	}
}
