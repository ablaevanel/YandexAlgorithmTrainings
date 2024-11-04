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
	scanner.Buffer(make([]byte, 0, 1024*1024*10), 1024*1024*10)
	scanner.Scan()
	n, _ := strconv.Atoi(strings.Trim(scanner.Text(), " "))
	arrA := make([][]int, n)
	arrB := make([][]int, n)
	a := make([]int, n)
	b := make([]int, n)
	p := make([]int, n)
	scanner.Scan()
	nums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(nums[i])

	}
	scanner.Scan()
	nums = strings.Split(strings.Trim(scanner.Text(), " "), " ")
	for i := 0; i < n; i++ {
		b[i], _ = strconv.Atoi(nums[i])
	}
	for i := 0; i < n; i++ {
		arrA[i] = []int{a[i], b[i], i + 1}
		arrB[i] = []int{b[i], a[i], i + 1}
	}
	scanner.Scan()
	nums = strings.Split(strings.Trim(scanner.Text(), " "), " ")
	for i := 0; i < n; i++ {
		p[i], _ = strconv.Atoi(nums[i])
	}
	sort.Slice(arrA, func(i, j int) bool {
		if arrA[i][0] == arrA[j][0] {
			if arrA[i][1] == arrA[j][1] {
				return arrA[i][2] < arrA[j][2]
			} else {
				return arrA[i][1] > arrA[j][1]
			}
		} else {
			return arrA[i][0] > arrA[j][0]
		}
	})
	sort.Slice(arrB, func(i, j int) bool {
		if arrB[i][0] == arrB[j][0] {
			if arrB[i][1] == arrB[j][1] {
				return arrB[i][2] < arrB[j][2]
			} else {
				return arrB[i][1] > arrB[j][1]
			}
		} else {
			return arrB[i][0] > arrB[j][0]
		}
	})
	output, _ := os.Create("output.txt")
	defer output.Close()
	isLearned := make([]bool, n)
	i, k, j := 0, 0, 0
	for i < n && k < n && j < n {
		if p[i] == 1 && k < n {
			for k < n && isLearned[arrB[k][2]-1] {
				k++
			}
			output.WriteString(fmt.Sprintf("%d ", arrB[k][2]))
			isLearned[arrB[k][2]-1] = true
			k++
		} else {
			for j < n && isLearned[arrA[j][2]-1] {
				j++
			}
			output.WriteString(fmt.Sprintf("%d ", arrA[j][2]))
			isLearned[arrA[j][2]-1] = true
			j++
		}
		i++
	}
}
