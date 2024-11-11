package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func isInt(s string) bool {
	flag := false
	if len(s) > 1 && (string(s[0]) == "+" || string(s[0]) == "-") {
		flag = true
	}
	for i, c := range s {
		if flag && i == 0 {
			continue
		}
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 10*1024*1024), 10*1024*1024)
	scanner.Scan()
	expression := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	stack := Stack{}
	rpn := []string{}
	priority := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
	}
	lastElem := ""
	for _, elem := range expression {
		if lastElem != "" && isInt(lastElem) && isInt(elem) && isInt(lastElem+elem) {
			fmt.Print("WRONG")
			return
		}
		lastElem = elem
	}
	word := ""
	line := strings.Join(expression, "")
	expression = []string{}
	for _, ch := range line {
		if ch >= '0' && ch <= '9' {
			word += string(ch)
		} else {
			if word != "" {
				expression = append(expression, word)
				word = ""
			}
			expression = append(expression, string(ch))
		}
	}
	if word != "" {
		expression = append(expression, word)
	}
	lastElem = ""
	for i, elem := range expression {
		if i == 0 && !(elem == "-" || isInt(elem) || elem == "(") {
			fmt.Print("WRONG")
			return
		}
		if i == len(expression)-1 && !(isInt(elem) || elem == ")") {
			fmt.Print("WRONG")
			return
		}
		if i == 0 && elem == "-" || lastElem == "(" && elem == "-" {
			rpn = append(rpn, "0")
		}
		_, ok1 := priority[lastElem]
		_, ok2 := priority[elem]
		if lastElem != "" && (lastElem == ")" && isInt(elem) || isInt(lastElem) && elem == "(" || lastElem == "(" && !isInt(elem) && !(elem == "-" || elem == "(") || ok1 && ok2) {
			fmt.Print("WRONG")
			return
		}
		if isInt(elem) {
			rpn = append(rpn, elem)
		} else if priority_value, ok := priority[elem]; ok {
			peeked, err := stack.peek()
			for err == nil && priority[peeked] >= priority_value {
				stack.pop()
				rpn = append(rpn, peeked)
				peeked, err = stack.peek()
			}
			stack.push(elem)
		} else if elem == "(" {
			stack.push(elem)
		} else if elem == ")" {
			popped, err := stack.pop()
			for err == nil && popped != "(" {
				rpn = append(rpn, popped)
				popped, err = stack.pop()
			}
			if err != nil {
				fmt.Print("WRONG")
				return
			}
		} else {
			fmt.Print("WRONG")
			return
		}
		lastElem = elem
	}
	for !stack.isEmpty() {
		popped, _ := stack.pop()
		if popped == "(" {
			fmt.Print("WRONG")
			return
		}
		rpn = append(rpn, popped)
	}
	for _, elem := range rpn {
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
