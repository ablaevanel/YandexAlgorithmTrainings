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
	arr := make([]int64, n)
	scanner.Scan()
	nums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	prefixSum := make([]int64, n+1)
	const mod = 1_000_000_007
	for i := 1; i <= n; i++ {
		arr[i-1], _ = strconv.ParseInt(nums[i-1], 10, 64)
		prefixSum[i] = (prefixSum[i-1] + arr[i-1]) % mod
	}
	output, _ := os.Create("output.txt")
	defer output.Close()
	prefixMultiply := make([]int64, n-1)
	for i := 1; i < n-1; i++ {
		prefixMultiply[i] = (prefixMultiply[i-1] + arr[i]*((prefixSum[n]-prefixSum[i+1]+mod)%mod)) % mod
	}
	var sum int64
	for i := 0; i < n-2; i++ {
		sum += (arr[i] * ((prefixMultiply[n-2] - prefixMultiply[i] + mod) % mod)) % mod
		sum = sum % mod
	}
	output.WriteString(strconv.FormatInt(sum, 10))
}
