package main

import "fmt"

func main() {
	var n, k, m, cnt int
	fmt.Scan(&n, &k, &m)
	for n/k > 0 && k/m > 0 {
		cnt += (n / k) * (k / m)
		n = (n % k) + (n/k)*(k%m)
	}
	fmt.Print(cnt)
}
