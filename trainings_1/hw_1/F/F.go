package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min4(a, b, c, d int) int {
	return min(min(a, b), min(c, d))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	v1 := max(a, c) * (b + d)
	v2 := max(a, d) * (b + c)
	v3 := max(b, c) * (a + d)
	v4 := max(b, d) * (a + c)
	switch min4(v1, v2, v3, v4) {
	case v1:
		fmt.Print(max(a, c), b+d)
	case v2:
		fmt.Print(max(a, d), b+c)
	case v3:
		fmt.Print(max(b, c), a+d)
	case v4:
		fmt.Print(max(b, d), a+c)
	}
}
