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
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Buffer(make([]byte, 0, 1024*1024*5), 1024*1024*5)
	scanner.Scan()
	n, _ := strconv.Atoi(strings.Trim(scanner.Text(), " "))
	arr := make([]int, n)
	scanner.Scan()
	nums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(nums[i])
	}
	sort.Ints(arr)
	var i, j int
	output, _ := os.Create("output.txt")
	defer output.Close()
	if n%2 == 0 {
		i, j = n/2-1, n/2
	} else {
		i, j = n/2, n/2
	}
	for n > 0 {
		if n%2 == 0 {
			output.WriteString(fmt.Sprintf("%d ", arr[i]))
			if i == j {
				j++
			}
			i--
		} else {
			output.WriteString(fmt.Sprintf("%d ", arr[j]))
			if i == j {
				i--
			}
			j++
		}
		n--
	}
}
