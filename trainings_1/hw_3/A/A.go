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
	arr := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	m := make(map[int]bool)
	var n int
	for _, str := range arr {
		n, _ = strconv.Atoi(str)
		m[n] = true
	}
	fmt.Print(len(m))
}
