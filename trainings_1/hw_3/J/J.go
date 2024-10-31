package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func extend(rect []int, delta int) {
	rect[0] -= delta
	rect[1] += delta
	rect[2] -= delta
	rect[3] += delta
}

func intersect(rect1, rect2 []int) {
	rect1[0] = max(rect1[0], rect2[0])
	rect1[1] = min(rect1[1], rect2[1])
	rect1[2] = max(rect1[2], rect2[2])
	rect1[3] = min(rect1[3], rect2[3])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	scanner.Scan()
	nums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	t, _ := strconv.Atoi(nums[0])
	d, _ := strconv.Atoi(nums[1])
	n, _ := strconv.Atoi(nums[2])
	rect := []int{0, 0, 0, 0}
	for i := 0; i < n; i++ {
		extend(rect, t)
		scanner.Scan()
		nav := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		x, _ := strconv.Atoi(nav[0])
		y, _ := strconv.Atoi(nav[1])
		navRect := []int{x - y, x - y, x + y, x + y}
		extend(navRect, d)
		intersect(rect, navRect)
	}
	var ans [][]int
	for i := rect[0]; i <= rect[1]; i++ {
		for j := rect[2]; j <= rect[3]; j++ {
			if (i+j)%2 == 0 {
				x := (i + j) / 2
				y := j - x
				ans = append(ans, []int{x, y})
			}
		}
	}
	fmt.Println(len(ans))
	for _, point := range ans {
		fmt.Println(point[0], point[1])
	}
}
