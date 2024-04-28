package main

import "fmt"

func main() {
	var nf, add string
	for n := 0; n < 4; n++ {
		fmt.Scanln(&nf)
		f := ""
		for i := range nf {
			if nf[i] >= '0' && nf[i] <= '9' {
				if i == 0 {
					if nf[i] != '8' {
						f += string(nf[i])
					}
				} else if i == 1 {
					if nf[0] != '+' {
						f += string(nf[i])
					}
				} else {
					f += string(nf[i])
				}
			}
		}
		if len(f) == 7 {
			f = "495" + f
		}
		if n == 0 {
			add = f
		} else {
			if f == add {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}
