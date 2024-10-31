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
	var x int
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		x, _ = strconv.Atoi(nums[0])
		m[x] = true
	}
	fmt.Print(len(m))
}
