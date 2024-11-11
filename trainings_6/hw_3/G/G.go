package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 10*1024*1024), 10*1024*1024)
	scanner.Scan()
	nums := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	n, _ := strconv.Atoi(nums[0])
	b, _ := strconv.Atoi(nums[1])
	scanner.Scan()
	arr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	intArr := make([]int, n)
	sum := 0
	odd := 0
	for i := 0; i < n; i++ {
		intArr[i], _ = strconv.Atoi(arr[i])
		sum += intArr[i] + odd
		currOdd := 0
		if odd >= b {
			currOdd = b
			odd -= b
		} else {
			currOdd = odd
			odd = 0
		}
		if intArr[i] > b-currOdd {
			odd += intArr[i] - b + currOdd
		}
	}
	sum += odd
	fmt.Print(sum)
}
