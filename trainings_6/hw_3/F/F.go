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
	scanner.Buffer(make([]byte, 0, 100000), 100000)
	scanner.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	order := strings.Split(strings.TrimSpace(scanner.Text()), "")
	priorityOfScope := make(map[rune]int)
	for i, scope := range order {
		priorityOfScope[rune(scope[0])] = i
	}
	scanner.Scan()
	ans := []rune(strings.TrimSpace(scanner.Text()))
	cnt := len(ans)
	stack := make([]rune, 0, 100000)
	pair := map[rune]rune{
		'(': ')',
		'[': ']',
	}
	for _, elem := range ans {
		if elem == ']' || elem == ')' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, elem)
		}
	}
	for cnt != n {
		if len(stack) == n-cnt {
			elem := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, pair[elem])
		} else {
			var elem rune
			if len(stack) != 0 {
				elem = stack[len(stack)-1]
			}
			if len(stack) != 0 && priorityOfScope[pair[elem]] < priorityOfScope['('] && priorityOfScope[pair[elem]] < priorityOfScope['['] {
				ans = append(ans, pair[elem])
				stack = stack[:len(stack)-1]
			} else {
				if priorityOfScope['('] < priorityOfScope['['] {
					stack = append(stack, '(')
					ans = append(ans, '(')
				} else {
					stack = append(stack, '[')
					ans = append(ans, '[')
				}
			}
		}
		cnt++
	}
	output, _ := os.Create("output.txt")
	defer output.Close()
	output.WriteString(string(ans))
}
