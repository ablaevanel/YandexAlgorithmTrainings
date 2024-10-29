package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var start, end float64 = 30, 4000
	var prev, curr float64
	var frequency string
	fmt.Scan(&prev)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&curr)
		fmt.Scan(&frequency)
		if math.Abs(curr-prev) < math.Pow(10, -6) {
			continue
		}
		if frequency == "closer" {
			if curr < prev {
				end = math.Min(end, (prev+curr)/2)
			} else {
				start = math.Max(start, (prev+curr)/2)
			}
		} else {
			if curr < prev {
				start = math.Max(start, (prev+curr)/2)
			} else {
				end = math.Min(end, (prev+curr)/2)
			}
		}
		prev = curr
	}
	fmt.Printf("%f %f", start, end)
}
