package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	var max, iMax int
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if i == 0 || arr[i] > max {
			max = arr[i]
			iMax = i
		}
	}
	meters := 0
	for i := 0; i < n; i++ {
		if i != 0 && i != n-1 && arr[i]%10 == 5 && arr[i+1] < arr[i] && arr[i] > meters && i > iMax {
			meters = arr[i]
		}
	}
	place := 1
	for i := 0; i < n; i++ {
		if arr[i] > meters {
			place++
		}
	}
	if meters == 0 {
		fmt.Print(0)
	} else {
		fmt.Print(place)
	}
}
