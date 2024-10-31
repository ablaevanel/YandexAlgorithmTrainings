package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	dict := make(map[string]bool)
	for scanner.Scan() {
		words := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		for _, word := range words {
			compare := ""
			for i := 0; i < len(word); i++ {
				compare += " "
			}
			if word == compare {
				continue
			}
			dict[word] = true
		}
	}
	output, _ := os.Create("output.txt")
	defer output.Close()
	output.WriteString(strconv.Itoa(len(dict)))
}
