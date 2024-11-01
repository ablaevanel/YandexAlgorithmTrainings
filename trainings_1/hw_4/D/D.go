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
	keys := make([]int, n)
	scanner.Scan()
	arr := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	for i, a := range arr {
		keys[i], _ = strconv.Atoi(a)
	}
	scanner.Scan()
	k, _ := strconv.Atoi(strings.Trim(scanner.Text(), " "))
	scanner.Scan()
	arr = strings.Split(strings.Trim(scanner.Text(), " "), " ")
	for i := 0; i < k; i++ {
		numOfKey, _ := strconv.Atoi(arr[i])
		keys[numOfKey-1]--
	}
	for _, v := range keys {
		if v < 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
