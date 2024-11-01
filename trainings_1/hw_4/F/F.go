package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	output, _ := os.Create("output.txt")
	defer output.Close()
	purchases := make(map[string]map[string]int)
	for scanner.Scan() {
		purchase := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		buyer := purchase[0]
		product := purchase[1]
		num, _ := strconv.Atoi(purchase[2])
		if purchases[buyer] == nil {
			purchases[buyer] = make(map[string]int)
		}
		purchases[buyer][product] += num
	}
	buyers := []string{}
	for buyer := range purchases {
		buyers = append(buyers, buyer)
	}
	sort.Strings(buyers)
	for _, buyer := range buyers {
		output.WriteString(fmt.Sprintf("%s:\n", buyer))
		products := []string{}
		for product := range purchases[buyer] {
			products = append(products, product)
		}
		sort.Strings(products)
		for _, product := range products {
			output.WriteString(fmt.Sprintf("%s %d\n", product, purchases[buyer][product]))
		}
	}
}
