package main

import (
	"bufio"
	"math"
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
	n, _ := strconv.ParseUint(strings.Trim(scanner.Text(), " "), 10, 64)
	arr := make([]uint64, n)
	revArr := make([]uint64, n)
	scanner.Scan()
	nums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	var i, k uint64
	prefSum := make([]uint64, n)
	prefOfPrefSum := make([]uint64, n)
	revPrefSum := make([]uint64, n)
	revPrefOfPrefSum := make([]uint64, n)
	for i < n {
		arr[i], _ = strconv.ParseUint(nums[i], 10, 64)
		revArr[n-i-1] = arr[i]
		i++
	}
	k = 1
	for k < n {
		prefSum[k] = prefSum[k-1] + arr[k-1]
		prefOfPrefSum[k] = prefOfPrefSum[k-1] + prefSum[k]
		k++
	}
	k = 1
	for k < n {
		revPrefSum[k] = revPrefSum[k-1] + revArr[k-1]
		revPrefOfPrefSum[k] = revPrefOfPrefSum[k-1] + revPrefSum[k]
		k++
	}
	output, _ := os.Create("output.txt")
	defer output.Close()
	var minSwaps uint64 = math.MaxUint64
	i = 0
	for i < n {
		var swaps uint64 = prefOfPrefSum[i] + revPrefOfPrefSum[n-i-1]
		if swaps < minSwaps && swaps != 0 {
			minSwaps = swaps
		}
		i++
	}
	output.WriteString(strconv.FormatUint(minSwaps, 10))
}
