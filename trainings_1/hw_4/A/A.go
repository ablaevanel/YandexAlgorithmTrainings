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
	m := make(map[string]string, n*2)
	var k, v string
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		k = arr[0]
		v = arr[1]
		m[k] = v
		m[v] = k
	}
	scanner.Scan()
	find := strings.Trim(scanner.Text(), " ")
	fmt.Print(m[find])
}
