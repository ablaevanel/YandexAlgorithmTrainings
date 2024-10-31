package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	scanner.Scan()
	arr1 := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	scanner.Scan()
	arr2 := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	m := make(map[string]int)
	for _, num := range arr1 {
		m[num]++
	}
	var ans []int
	for _, num := range arr2 {
		if m[num] > 0 {
			n, _ := strconv.Atoi(num)
			ans = append(ans, n)
			m[num]--
		}
	}
	sort.Ints(ans)
	output, err := os.Create("output.txt")
	if err != nil {
		os.Exit(1)
	}
	defer output.Close()
	for _, num := range ans {
		output.WriteString(strconv.Itoa(num) + " ")
	}
}
