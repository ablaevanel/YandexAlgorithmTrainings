package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	scanner.Scan()
	n, _ := strconv.Atoi(strings.Trim(scanner.Text(), " "))
	blocks := make(map[int]int)
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		if blocks[x] < y {
			blocks[x] = y
		}
	}
	ans := 0
	for _, v := range blocks {
		ans += v
	}
	fmt.Print(ans)
}
