package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	var answers [][]int
	if a > 0 && c > 0 {
		answers = append(answers, []int{b + 1, d + 1})
	}
	if b > 0 && d > 0 {
		answers = append(answers, []int{a + 1, c + 1})
	}
	if a > 0 && b > 0 {
		answers = append(answers, []int{max(a, b) + 1, 1})
	}
	if c > 0 && d > 0 {
		answers = append(answers, []int{1, max(c, d) + 1})
	}
	m, n := answers[0][0], answers[0][1]
	for _, ans := range answers {
		if ans[0]+ans[1] < m+n {
			m = ans[0]
			n = ans[1]
		}
	}
	fmt.Print(m, n)
}
