package main

import "fmt"

func main() {
	var a, b, c, d, e, f float64
	fmt.Scan(&a, &b, &c, &d, &e, &f)
	if (a*d - b*c) != 0 {
		fmt.Print(2, (e*d-b*f)/(a*d-b*c), (a*f-e*c)/(a*d-b*c))
	} else {
		if (e*d-b*f) == 0 && (a*f-e*c) == 0 {
			if a == 0 && b == 0 && c == 0 && d == 0 {
				if e == 0 && f == 0 {
					fmt.Print(5)
				} else {
					fmt.Print(0)
				}
			} else if b == 0 && d == 0 {
				if a != 0 {
					fmt.Print(3, e/a)
				} else {
					fmt.Print(3, f/c)
				}
			} else if a == 0 && c == 0 {
				if b != 0 {
					fmt.Print(4, e/b)
				} else {
					fmt.Print(4, f/d)
				}
			} else if b != 0 || d != 0 {
				if b != 0 {
					fmt.Print(1, -a/b, e/b)
				} else {
					fmt.Print(1, -c/d, f/d)
				}
			}
		} else {
			fmt.Print(0)
		}
	}
}
