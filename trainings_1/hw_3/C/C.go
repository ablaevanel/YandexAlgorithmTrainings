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
	var n, m int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	scanner.Scan()
	nums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	n, _ = strconv.Atoi(nums[0])
	m, _ = strconv.Atoi(nums[1])
	a := make(map[int]bool, n)
	b := make(map[int]bool, m)
	ans := make([]int, 0, n+m)
	ansAnya := make([]int, 0, n)
	ansBorya := make([]int, 0, m)
	var num int
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ = strconv.Atoi(scanner.Text())
		a[num] = true
	}
	for i := 0; i < m; i++ {
		scanner.Scan()
		num, _ = strconv.Atoi(scanner.Text())
		b[num] = true
		if a[num] {
			ans = append(ans, num)
		} else {
			ansBorya = append(ansBorya, num)
		}
	}
	for k := range a {
		if !b[k] {
			ansAnya = append(ansAnya, k)
		}
	}
	sort.Ints(ans)
	fmt.Println(len(ans))
	for i := 0; i < len(ans); i++ {
		fmt.Printf("%d ", ans[i])
	}
	fmt.Println()
	sort.Ints(ansAnya)
	fmt.Println(len(ansAnya))
	for i := 0; i < len(ansAnya); i++ {
		fmt.Printf("%d ", ansAnya[i])
	}
	fmt.Println()
	sort.Ints(ansBorya)
	fmt.Println(len(ansBorya))
	for i := 0; i < len(ansBorya); i++ {
		fmt.Printf("%d ", ansBorya[i])
	}
}
