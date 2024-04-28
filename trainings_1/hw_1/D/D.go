package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if c < 0 {
		fmt.Print("NO SOLUTION")
	} else {
		if a == 0 {
			if c*c-b == 0 {
				fmt.Print("MANY SOLUTIONS")
			} else {
				fmt.Print("NO SOLUTION")
			}
		} else {
			if (c*c-b)%a != 0 {
				fmt.Print("NO SOLUTION")
			} else {
				fmt.Print((c*c - b) / a)
			}
		}
	}
}
