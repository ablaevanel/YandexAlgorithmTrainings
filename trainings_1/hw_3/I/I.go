package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		for j := 0; j < k; j++ {
			scanner.Scan()
			m[scanner.Text()]++
		}
	}
	var all, oneOrMore []string
	for k, v := range m {
		if v == n {
			all = append(all, k)
		}
		oneOrMore = append(oneOrMore, k)
	}
	fmt.Println(len(all))
	for i := 0; i < len(all); i++ {
		fmt.Println(all[i])
	}
	fmt.Println(len(oneOrMore))
	for i := 0; i < len(oneOrMore); i++ {
		fmt.Println(oneOrMore[i])
	}
}
