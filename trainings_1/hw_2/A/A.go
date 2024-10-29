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
	s := scanner.Text()
	arr := strings.Split(s, " ")
	intArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		intArr[i], _ = strconv.Atoi(arr[i])
	}
	for i := 1; i < len(intArr); i++ {
		if intArr[i] <= intArr[i-1] {
			fmt.Print("NO")
			return
		}
	}
	fmt.Print("YES")
}
