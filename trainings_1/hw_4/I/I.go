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
	scanner.Buffer(make([]byte, 0, 1024*1024*5), 1024*1024*5)
	scanner.Scan()
	n, _ := strconv.Atoi(strings.Trim(scanner.Text(), " "))
	dict := make(map[string][]string)
	for i := 0; i < n; i++ {
		scanner.Scan()
		word := strings.Trim(scanner.Text(), " ")
		key := strings.ToLower(word)
		dict[key] = append(dict[key], word)
	}
	scanner.Scan()
	words := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	cnt := 0
	for _, word := range words {
		compare := ""
		for i := 0; i < len(word); i++ {
			compare += " "
		}
		if compare == word {
			continue
		}
		key := strings.ToLower(word)
		if variants, ok := dict[key]; ok {
			flag := true
			for _, variant := range variants {
				if variant == word {
					flag = false
				}
			}
			if flag {
				cnt++
			}
		} else {
			cntOfUpper := 0
			for i := 0; i < len(word); i++ {
				if word[i] >= 'A' && word[i] <= 'Z' {
					cntOfUpper++
				}
			}
			if cntOfUpper != 1 {
				cnt++
			}
		}
	}
	output, _ := os.Create("output.txt")
	defer output.Close()
	output.WriteString(strconv.Itoa(cnt))
}
