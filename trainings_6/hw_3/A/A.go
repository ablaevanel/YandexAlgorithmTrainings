package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Stack []rune

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) pop() (rune, error) {
	if s.isEmpty() {
		return 0, errors.New("Stack is empty")
	}
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped, nil
}

func (s *Stack) peek() (rune, error) {
	if s.isEmpty() {
		return 0, errors.New("Stack is empty")
	}
	peek := (*s)[len(*s)-1]
	return peek, nil
}

func (s *Stack) push(elem rune) {
	*s = append((*s), elem)
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 10*1024*1024), 10*1024*1024)
	scanner.Scan()
	text := scanner.Text()
	stack := Stack{}
	for _, ch := range text {
		peek, _ := stack.peek()
		switch ch {
		case '(':
			stack.push(ch)
		case '[':
			stack.push(ch)
		case '{':
			stack.push(ch)
		case '}':
			if stack.isEmpty() || peek != '{' {
				fmt.Print("no")
				return
			}
			stack.pop()
		case ']':
			if stack.isEmpty() || peek != '[' {
				fmt.Print("no")
				return
			}
			stack.pop()
		case ')':
			if stack.isEmpty() || peek != '(' {
				fmt.Print("no")
				return
			}
			stack.pop()
		}
	}
	if stack.isEmpty() {
		fmt.Print("yes")
	} else {
		fmt.Print("no")
	}
}
