package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	if d > e {
		d, e = e, d
	}
	s := a + b + c
	s1 := min(min(a, b), c)
	s2 := s - s1 - max(max(a, b), c)
	if d >= s1 && e >= s2 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
