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
	scanner.Buffer(make([]byte, 0, 10*1024*1024), 10*1024*1024)
	scanner.Scan()
	nums := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	n, _ := strconv.Atoi(nums[0])
	H, _ := strconv.Atoi(nums[1])
	hw := make([][]int, n)
	scanner.Scan()
	nums = strings.Split(strings.TrimSpace(scanner.Text()), " ")
	for i, num := range nums {
		height, _ := strconv.Atoi(num)
		hw[i] = []int{height}
	}
	scanner.Scan()
	nums = strings.Split(strings.TrimSpace(scanner.Text()), " ")
	for i, num := range nums {
		width, _ := strconv.Atoi(num)
		hw[i] = append(hw[i], width)
	}
	sort.Slice(hw, func(i, j int) bool {
		return hw[i][0] < hw[j][0]
	})
	deque := []int{}
	i, j := 0, 1
	currH := hw[0][1]
	maxDiff := -1
	if currH >= H {
		fmt.Print(0)
		return
	}
	for !(j == n && currH < H) {
		if j < n && currH < H {
			for len(deque) != 0 && deque[len(deque)-1] < hw[j][0]-hw[j-1][0] {
				deque = deque[:len(deque)-1]
			}
			deque = append(deque, hw[j][0]-hw[j-1][0])
			currH += hw[j][1]
			j++
		} else {
			if i < n && hw[i][1] >= H {
				fmt.Print(0)
				return
			}
			if len(deque) != 0 && (maxDiff == -1 || maxDiff > deque[0]) {
				maxDiff = deque[0]
			}
			if len(deque) != 0 && i < n-1 && deque[0] == hw[i+1][0]-hw[i][0] {
				deque = deque[1:]
			}
			currH -= hw[i][1]
			i++
		}
	}
	fmt.Print(maxDiff)
}
