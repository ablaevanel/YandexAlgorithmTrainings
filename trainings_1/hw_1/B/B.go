package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a+b > c && b+c > a && a+c > b {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
