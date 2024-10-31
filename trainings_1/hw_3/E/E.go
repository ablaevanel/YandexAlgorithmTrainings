package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	var n int
	fmt.Scan(&n)
	nums := make([]bool, 10)
	nums[x] = true
	nums[y] = true
	nums[z] = true
	cnt := 0
	for n > 0 {
		odd := n % 10
		if !nums[odd] {
			nums[odd] = true
			cnt++
		}
		n /= 10
	}
	fmt.Print(cnt)
}
