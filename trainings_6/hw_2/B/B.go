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
	k, _ := strconv.Atoi(nums[1])
	scanner.Scan()
	nums = strings.Split(strings.Trim(scanner.Text(), " "), " ")
	numArr := make([]int, n)
	prefix := make([]int, n+1)
	i, j := 0, 0
	for i, num := range nums {
		number, _ := strconv.Atoi(num)
		numArr[i] = number
		prefix[i+1] = prefix[i] + numArr[i]
	}
	cnt := 0
	for i <= j && j <= n && i <= n {
		if prefix[j]-prefix[i] == k {
			cnt++
			i++
			j++
		} else if prefix[j]-prefix[i] < k {
			j++
		} else {
			i++
		}
	}
	ouput, _ := os.Create("output.txt")
	defer ouput.Close()
	ouput.WriteString(strconv.Itoa(cnt))
}
