package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	words := make(map[string]int)
	output, _ := os.Create("output.txt")
	defer output.Close()
	for scanner.Scan() {
		line := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		for _, word := range line {
			compare := ""
			for i := 0; i < len(word); i++ {
				compare += " "
			}
			if compare == word {
				continue
			}
			words[word]++
		}
	}
	ans := ""
	cnt := 0
	for k, v := range words {
		if v > cnt || v == cnt && (ans == "" || ans > k) {
			cnt = v
			ans = k
		}
	}
	output.WriteString(ans)
}
