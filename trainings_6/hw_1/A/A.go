package main

import (
	"fmt"
)

func main() {
	var x1, y1, x2, y2, x, y float64
	fmt.Scan(&x1, &y1, &x2, &y2, &x, &y)
	if y > y2 {
		fmt.Print("N")
	}
	if y < y1 {
		fmt.Print("S")
	}
	if x < x1 {
		fmt.Print("W")
	}
	if x > x2 {
		fmt.Print("E")
	}
}
