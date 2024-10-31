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
	n, _ := strconv.Atoi(scanner.Text())
	m := make(map[int]bool)
	var a, b int
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		a, _ = strconv.Atoi(nums[0])
		b, _ = strconv.Atoi(nums[1])
		if a+b == n-1 && a >= 0 && b >= 0 {
			m[a] = true
		}
	}
	fmt.Print(len(m))
}
