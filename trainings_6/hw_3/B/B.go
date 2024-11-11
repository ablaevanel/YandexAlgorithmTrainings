package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack [][]int

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) pop() ([]int, error) {
	if s.isEmpty() {
		return []int{}, errors.New("Stack is empty")
	}
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped, nil
}

func (s *Stack) peek() ([]int, error) {
	if s.isEmpty() {
		return []int{}, errors.New("Stack is empty")
	}
	peek := (*s)[len(*s)-1]
	return peek, nil
}

func (s *Stack) push(elem []int) {
	*s = append((*s), elem)
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 10*1024*1024), 10*1024*1024)
	scanner.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	arr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	intArr := make([]int, n)
	for i, num := range arr {
		intArr[i], _ = strconv.Atoi(num)
	}
	stack := Stack{}
	ans := make([]int, n)
	for i, num := range intArr {
		peeked, err := stack.peek()
		for err == nil && peeked[0] > num {
			ans[peeked[1]] = i
			stack.pop()
			peeked, err = stack.peek()
		}
		stack.push([]int{num, i})
	}
	for !stack.isEmpty() {
		popped, _ := stack.pop()
		ans[popped[1]] = -1
	}
	for _, num := range ans {
		fmt.Printf("%d ", num)
	}
}
