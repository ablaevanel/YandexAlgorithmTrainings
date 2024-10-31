package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	scanner.Scan()
	firstGenom := scanner.Text()
	scanner.Scan()
	secondGenom := scanner.Text()
	m := make(map[string]bool, len(secondGenom)-1)
	for i := 0; i < len(secondGenom)-1; i++ {
		gen := string(secondGenom[i]) + string(secondGenom[i+1])
		m[gen] = true
	}
	cnt := 0
	for i := 0; i < len(firstGenom)-1; i++ {
		gen := string(firstGenom[i]) + string(firstGenom[i+1])
		if m[gen] {
			cnt++
		}
	}
	fmt.Print(cnt)
}
