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
	var a, b, na, nb int
	fmt.Scan(&a, &b, &na, &nb)
	mina := a*(na-1) + na
	maxa := a*(na+1) + na
	minb := b*(nb-1) + nb
	maxb := b*(nb+1) + nb
	if mina > maxb || minb > maxa {
		fmt.Print(-1)
	} else {
		fmt.Print(max(mina, minb), min(maxa, maxb))
	}
}
