package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []int

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("Stack is empty")
	}
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped, nil
}

func (s *Stack) push(elem int) {
	*s = append((*s), elem)
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 10*1024*1024), 10*1024*1024)
	scanner.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	stack := Stack{}
	prefSum := []int{0}
	for i := 0; i < n; i++ {
		scanner.Scan()
		operation := strings.TrimSpace(scanner.Text())
		if operation[0] == '+' {
			add, _ := strconv.Atoi(operation[1:])
			stack.push(add)
			prefSum = append(prefSum, prefSum[len(prefSum)-1]+add)
		} else if operation == "-" {
			elem, _ := stack.pop()
			prefSum = prefSum[:len(prefSum)-1]
			fmt.Println(elem)
		} else if operation[0] == '?' {
			k, _ := strconv.Atoi(operation[1:])
			sum := prefSum[len(prefSum)-1] - prefSum[len(prefSum)-k-1]
			fmt.Println(sum)
		}
	}
}
