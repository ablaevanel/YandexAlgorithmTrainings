package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Buffer(make([]byte, 0, 1024*1024*10), 1024*1024*10)
	scanner.Scan()
	nums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	n, _ := strconv.Atoi(nums[0])
	k, _ := strconv.ParseInt(nums[1], 10, 64)
	arr := make([]int64, n)
	scanner.Scan()
	nums = strings.Split(strings.Trim(scanner.Text(), " "), " ")
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.ParseInt(nums[i], 10, 64)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	output, _ := os.Create("output.txt")
	defer output.Close()
	cnt := 1
	i, j := 0, 1
	for i < j && j < n {
		for j < n && arr[j]-arr[i] <= k {
			j++
		}
		if cnt < j-i {
			cnt = j - i
		}
		i++
		if i == j {
			j++
		}
	}
	output.WriteString(strconv.Itoa(cnt))
}
