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
	scanner.Buffer(make([]byte, 0, 1024*1024*5), 1024*1024*5)
	scanner.Scan()
	nums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	n, _ := strconv.Atoi(nums[0])
	r, _ := strconv.Atoi(nums[1])
	arr := make([]int, n)
	scanner.Scan()
	nums = strings.Split(strings.Trim(scanner.Text(), " "), " ")
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(nums[i])
	}
	cnt := 0
	i, j := 0, 1
	for i <= j && j < n && i < n {
		if arr[j]-arr[i] > r {
			cnt += n - j
			i++
		} else {
			j++
		}
	}
	output, _ := os.Create("output.txt")
	defer output.Close()
	output.WriteString(strconv.Itoa(cnt))
}
