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
	k, _ := strconv.Atoi(nums[1])
	deque := make([]int, k)
	start, end := 0, 0
	scanner.Scan()
	arr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	intArr := make([]int, n)
	for i, num := range arr {
		intArr[i], _ = strconv.Atoi(num)
	}
	for i := 0; i < k; i++ {
		for (end-start+k)%k > 0 && deque[(end-1+k)%k] > intArr[i] {
			end = (end + k - 1) % k
		}
		deque[end] = intArr[i]
		end = (end + 1) % k
	}
	fmt.Println(deque[start])
	if intArr[0] == deque[start] {
		start = (start + 1) % k
	}
	for i := k; i < n; i++ {
		for (end-start+k)%k > 0 && deque[(end-1+k)%k] > intArr[i] {
			end = (end + k - 1) % k
		}
		deque[end] = intArr[i]
		end = (end + 1) % k
		fmt.Println(deque[start])
		if intArr[i-k+1] == deque[start] {
			start = (start + 1) % k
		}
	}
}
