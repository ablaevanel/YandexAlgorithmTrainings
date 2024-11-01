package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Buffer(make([]byte, 0, 1024*1024*5), 1024*1024*5)
	scanner.Scan()
	arr := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	n, _ := strconv.Atoi(arr[0])
	sensitiveToRegistry := arr[1]
	beginWithNumber := arr[2]
	keyWords := make(map[string]bool)
	for i := 0; i < n; i++ {
		scanner.Scan()
		keyWord := strings.Trim(scanner.Text(), " ")
		if sensitiveToRegistry == "no" {
			keyWord = strings.ToUpper(keyWord)
		}
		keyWords[keyWord] = true
	}
	identifires := make(map[string][]int)
	j := 0
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		n := len(line)
		word := ""
		i := -1
		for k, ch := range line {
			if unicode.IsDigit(ch) || unicode.Is(unicode.Latin, ch) || ch == '_' {
				word += string(ch)
				if i == -1 {
					i = k
				}
			} else {
				if len(word) == 0 || len(word) == 1 && unicode.IsDigit(rune(word[0])) {
					word = ""
					i = -1
					continue
				}
				if sensitiveToRegistry == "no" {
					word = strings.ToUpper(word)
				}
				if !keyWords[word] && (beginWithNumber == "no" && !unicode.IsDigit(rune(word[0])) || beginWithNumber == "yes") {
					if val, ok := identifires[word]; ok {
						val[1]++
					} else {
						identifires[word] = []int{i + j, 1}
					}
				}
				word = ""
				i = -1
			}
		}
		if len(word) == 0 || len(word) == 1 && unicode.IsDigit(rune(word[0])) {
			continue
		}
		if sensitiveToRegistry == "no" {
			word = strings.ToUpper(word)
		}
		if !keyWords[word] && (beginWithNumber == "no" && !unicode.IsDigit(rune(word[0])) || beginWithNumber == "yes") {
			if val, ok := identifires[word]; ok {
				val[1]++
			} else {
				identifires[word] = []int{i + j, 1}
			}
		}
		j += n
	}
	pos := -1
	ans := ""
	cnt := 0
	for k, v := range identifires {
		if v[1] > cnt {
			ans = k
			pos = v[0]
			cnt = v[1]
		} else if v[1] == cnt && v[0] < pos {
			ans = k
			pos = v[0]
		}
	}
	fmt.Print(ans)
}
