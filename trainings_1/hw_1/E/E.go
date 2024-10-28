package main

import (
	"fmt"
)

func main() {
	var k1, m, k2, p2, n2 int
	fmt.Scan(&k1, &m, &k2, &p2, &n2)
	if !(0 <= n2 && n2 <= m) {
		fmt.Print(-1, -1)
		return
	}
	var q []int
	if p2 == 1 && n2 == 1 {
		if k1 <= k2 {
			fmt.Print(1, 1)
			return
		} else {
			for i := k2; i <= k1; i++ {
				q = append(q, i)
			}
		}
	} else {
		minBound := k2 / (m*(p2-1) + n2)
		maxBound := (k2 - 1) / (m*(p2-1) + n2 - 1)
		for i := minBound; i <= maxBound; i++ {
			if i != 0 && ((m*(p2-1)+n2-1)*i+(k2-1)%i) == k2-1 {
				q = append(q, i)
			}
		}
	}
	p, n := -1, -1
	for _, value := range q {
		floorIndex := (k1-1-(k1-1)%value)/value + 1
		n1 := floorIndex % m
		p1 := (floorIndex-n1)/m + 1
		if n1 == 0 {
			n1 = m
			p1 -= 1
		}
		if p == -1 && n == -1 {
			p = p1
			n = n1
		} else {
			if p != p1 {
				p = 0
			}
			if n != n1 {
				n = 0
			}
		}
	}
	fmt.Print(p, n)
}
