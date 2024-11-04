package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Buffer(make([]byte, 0, 1024*1024*10), 1024*1024*10)
	scanner.Scan()
	n, _ := strconv.Atoi(strings.Trim(scanner.Text(), " "))
	scanner.Scan()
	weights := make([]int, n)
	nums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	for i, num := range nums {
		weights[i], _ = strconv.Atoi(num)
	}
	scanner.Scan()
	nums = strings.Split(strings.Trim(scanner.Text(), " "), " ")
	m, _ := strconv.Atoi(nums[0])
	k, _ := strconv.Atoi(nums[1])
	scanner.Scan()
	indexes := make([]int, m)
	nums = strings.Split(strings.Trim(scanner.Text(), " "), " ")
	for i, num := range nums {
		indexes[i], _ = strconv.Atoi(num)
	}
	answers := make([]int, n)
	j := n - 1
	for i := n - 1; i >= 0 && j >= 0 && i >= j; i-- {
		for j > 0 && (weights[j-1] < weights[j] || k > 0 && weights[j-1] == weights[j]) {
			if weights[j-1] == weights[j] {
				k--
			}
			j--
		}
		answers[i] = j + 1
		if i > 0 && j < i && weights[i-1] == weights[i] {
			k++
		}
		if j > 0 && i > 0 && i <= j {
			j = i - 1
		}
	}
	output, _ := os.Create("output.txt")
	defer output.Close()
	for i := 0; i < m; i++ {
		output.WriteString(strconv.Itoa(answers[indexes[i]-1]) + " ")
	}
}
