package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	bank := make(map[string]int)
	for scanner.Scan() {
		operation := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		switch operation[0] {
		case "DEPOSIT":
			sum, _ := strconv.Atoi(operation[2])
			bank[operation[1]] += sum
		case "WITHDRAW":
			sum, _ := strconv.Atoi(operation[2])
			bank[operation[1]] -= sum
		case "BALANCE":
			if balance, ok := bank[operation[1]]; ok {
				fmt.Println(balance)
			} else {
				fmt.Println("ERROR")
			}
		case "TRANSFER":
			sum, _ := strconv.Atoi(operation[3])
			bank[operation[1]] -= sum
			bank[operation[2]] += sum
		case "INCOME":
			percent, _ := strconv.Atoi(operation[1])
			for client := range bank {
				if bank[client] > 0 {
					bank[client] = bank[client] * (100 + percent) / 100
				}
			}
		}
	}
}
