package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := strings.Split(scanner.Text(), " ")
	floatArr := make([]float64, len(arr))
	for i := 0; i < len(arr); i++ {
		floatArr[i], _ = strconv.ParseFloat(arr[i], 64)
	}
	var cnt int
	for i := 1; i < len(floatArr)-1; i++ {
		if floatArr[i] > floatArr[i-1] && floatArr[i] > floatArr[i+1] {
			cnt++
		}
	}
	fmt.Print(cnt)
}
