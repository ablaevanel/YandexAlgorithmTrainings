package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) pop() (string, error) {
	if s.isEmpty() {
		return "", errors.New("Stack is empty")
	}
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped, nil
}

func (s *Stack) peek() (string, error) {
	if s.isEmpty() {
		return "", errors.New("Stack is empty")
	}
	peek := (*s)[len(*s)-1]
	return peek, nil
}

func (s *Stack) push(elem string) {
	*s = append((*s), elem)
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 10*1024*1024), 10*1024*1024)
	scanner.Scan()
	expression := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	stack := Stack{}
	for _, elem := range expression {
		if elem == "+" || elem == "-" || elem == "*" {
			elem1, _ := stack.pop()
			elem2, _ := stack.pop()
			num1, _ := strconv.Atoi(elem1)
			num2, _ := strconv.Atoi(elem2)
			num3 := 0
			switch elem {
			case "+":
				num3 = num2 + num1
			case "-":
				num3 = num2 - num1
			case "*":
				num3 = num2 * num1
			}
			stack.push(strconv.Itoa(num3))
		} else {
			stack.push(elem)
		}
	}
	ans, _ := stack.peek()
	fmt.Print(ans)
}
