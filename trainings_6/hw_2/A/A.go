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
	n, _ := strconv.Atoi(strings.Trim(scanner.Text(), " "))
	scanner.Scan()
	nums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	numArr := make([]int, n)
	prefix := make([]int, n+1)
	for i, num := range nums {
		number, _ := strconv.Atoi(num)
		numArr[i] = number
		prefix[i+1] = prefix[i] + numArr[i]
	}
	ouput, _ := os.Create("output.txt")
	defer ouput.Close()
	for i := 1; i <= n; i++ {
		ouput.WriteString(strconv.Itoa(prefix[i]) + " ")
	}
}
