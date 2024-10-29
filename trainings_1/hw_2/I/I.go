package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	field := make([][]string, n+2)
	for i := 0; i < n+2; i++ {
		field[i] = make([]string, m+2)
	}
	var x, y int
	for i := 0; i < k; i++ {
		fmt.Scan(&y, &x)
		field[y][x] = "*"
	}
	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{1, 1, 1, 0, 0, -1, -1, -1}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if field[i][j] == "*" {
				continue
			}
			cnt := 0
			for k := 0; k < len(dx); k++ {
				if field[i+dy[k]][j+dx[k]] == "*" {
					cnt++
				}
			}
			field[i][j] = strconv.Itoa(cnt)
		}
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			fmt.Print(field[i][j], " ")
		}
		fmt.Println()
	}
}
