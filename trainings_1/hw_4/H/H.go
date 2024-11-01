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
	scanner.Buffer(make([]byte, 0, 1024*1024*5), 1024*1024*5)
	scanner.Scan()
	lenNums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	len1, _ := strconv.Atoi(lenNums[0])
	len2, _ := strconv.Atoi(lenNums[1])

	scanner.Scan()
	word1 := strings.Trim(scanner.Text(), " ")
	letters := make(map[byte]int, len1)
	for i := range word1 {
		letters[word1[i]]++
	}
	scanner.Scan()
	word2 := strings.Trim(scanner.Text(), " ")
	i, j, cnt := 0, 0, 0
	for i <= j && j < len2 && i < len2 {
		if val, ok := letters[word2[j]]; ok {
			if val != 0 {
				letters[word2[j]]--
				j++
				if j-i == len1 {
					cnt++
				}
			} else {
				if _, ok := letters[word2[i]]; ok {
					letters[word2[i]]++
				}
				i++
			}
		} else {
			for i < j {
				if _, ok := letters[word2[i]]; ok {
					letters[word2[i]]++
				}
				i++
			}
			i++
			j++
		}
	}
	output, _ := os.Create("output.txt")
	defer output.Close()
	output.WriteString(strconv.Itoa(cnt))
}
