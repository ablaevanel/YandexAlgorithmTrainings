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
	scanner.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	nums := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	a, _ := strconv.Atoi(nums[0])
	b, _ := strconv.Atoi(nums[1])
	if a > b {
		a, b = b, a
	}
	d1 := make([][]int, 0, n)
	d2 := make([][]int, 0, n)
	d3 := make([][]int, 0, n)
	d4 := make([][]int, 0, n)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums = strings.Split(strings.TrimSpace(scanner.Text()), " ")
		d, _ := strconv.Atoi(nums[0])
		t, _ := strconv.Atoi(nums[1])
		switch d {
		case 1:
			d1 = append(d1, []int{t, i})
		case 2:
			d2 = append(d2, []int{t, i})
		case 3:
			d3 = append(d3, []int{t, i})
		case 4:
			d4 = append(d4, []int{t, i})
		}
	}
	sort.Slice(d1, func(i, j int) bool {
		return d1[i][0] < d1[j][0]
	})
	sort.Slice(d2, func(i, j int) bool {
		return d2[i][0] < d2[j][0]
	})
	sort.Slice(d3, func(i, j int) bool {
		return d3[i][0] < d3[j][0]
	})
	sort.Slice(d4, func(i, j int) bool {
		return d4[i][0] < d4[j][0]
	})
	p1, p2, p3, p4 := 0, 0, 0, 0
	i := 1
	for p1 != len(d1) || p3 != len(d3) || p2 != len(d2) || p4 != len(d4) {
		if a == 1 && b == 3 {
			if p1 != len(d1) && p3 != len(d3) && d1[p1][0] == i && d3[p3][0] == i {
				ans[d1[p1][1]] = i
				ans[d3[p3][1]] = i
				p1++
				p3++
			} else if p1 != len(d1) && d1[p1][0] == i {
				ans[d1[p1][1]] = i
				p1++
			} else if p3 != len(d3) && d3[p3][0] == i {
				ans[d3[p3][1]] = i
				p3++
			} else if p2 != len(d2) && p4 != len(d4) && d2[p2][0] <= i && d4[p4][0] <= i {
				ans[d2[p2][1]] = i
				ans[d4[p4][1]] = i
				p2++
				p4++
			} else if p2 != len(d2) && d2[p2][0] <= i {
				ans[d2[p2][1]] = i
				p2++
			} else if p4 != len(d4) && d4[p4][0] <= i {
				ans[d4[p4][1]] = i
				p4++
			}
		} else if a == 2 && b == 4 {
			if p2 != len(d2) && p4 != len(d4) && d2[p2][0] == i && d4[p4][0] == i {
				ans[d2[p2][1]] = i
				ans[d4[p4][1]] = i
				p2++
				p4++
			} else if p2 != len(d2) && d2[p2][0] == i {
				ans[d2[p2][1]] = i
				p2++
			} else if p4 != len(d4) && d4[p4][0] == i {
				ans[d4[p4][1]] = i
				p4++
			} else if p1 != len(d1) && p3 != len(d3) && d1[p1][0] <= i && d3[p3][0] <= i {
				ans[d1[p1][1]] = i
				ans[d3[p3][1]] = i
				p1++
				p3++
			} else if p1 != len(d1) && d1[p1][0] <= i {
				ans[d1[p1][1]] = i
				p1++
			} else if p3 != len(d3) && d3[p3][0] <= i {
				ans[d3[p3][1]] = i
				p3++
			}
		} else if a == 1 && b == 2 {
			if p1 != len(d1) && d1[p1][0] == i {
				ans[d1[p1][1]] = i
				p1++
			} else if p2 != len(d2) && d2[p2][0] <= i {
				ans[d2[p2][1]] = i
				p2++
			} else if p3 != len(d3) && d3[p3][0] <= i {
				ans[d3[p3][1]] = i
				p3++
			} else if p4 != len(d4) && d4[p4][0] <= i {
				ans[d4[p4][1]] = i
				p4++
			}
		} else if a == 2 && b == 3 {
			if p2 != len(d2) && d2[p2][0] == i {
				ans[d2[p2][1]] = i
				p2++
			} else if p3 != len(d3) && d3[p3][0] <= i {
				ans[d3[p3][1]] = i
				p3++
			} else if p4 != len(d4) && d4[p4][0] <= i {
				ans[d4[p4][1]] = i
				p4++
			} else if p1 != len(d1) && d1[p1][0] <= i {
				ans[d1[p1][1]] = i
				p1++
			}
		} else if a == 3 && b == 4 {
			if p3 != len(d3) && d3[p3][0] == i {
				ans[d3[p3][1]] = i
				p3++
			} else if p4 != len(d4) && d4[p4][0] <= i {
				ans[d4[p4][1]] = i
				p4++
			} else if p1 != len(d1) && d1[p1][0] <= i {
				ans[d1[p1][1]] = i
				p1++
			} else if p2 != len(d2) && d2[p2][0] <= i {
				ans[d2[p2][1]] = i
				p2++
			}
		} else if a == 1 && b == 4 {
			if p4 != len(d4) && d4[p4][0] == i {
				ans[d4[p4][1]] = i
				p4++
			} else if p1 != len(d1) && d1[p1][0] <= i {
				ans[d1[p1][1]] = i
				p1++
			} else if p2 != len(d2) && d2[p2][0] <= i {
				ans[d2[p2][1]] = i
				p2++
			} else if p3 != len(d3) && d3[p3][0] <= i {
				ans[d3[p3][1]] = i
				p3++
			}
		}
		i++
	}
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}
