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
	const maxCapacity = 1024 * 1024
	scanner.Buffer(make([]byte, 0, maxCapacity), maxCapacity)
	scanner.Scan()
	arr := strings.Split(scanner.Text(), " ")
	intArr := make([]int, len(arr))
	for i, str := range arr {
		intArr[i], _ = strconv.Atoi(str)
	}
	fn, sn := intArr[0], intArr[1]
	if fn > sn {
		fn, sn = sn, fn
	}
	sp, fp := fn, sn
	for i := 2; i < len(intArr); i++ {
		if intArr[i] > fp {
			sp = fp
			fp = intArr[i]
		} else if intArr[i] > sp {
			sp = intArr[i]
		}
		if intArr[i] < fn {
			sn = fn
			fn = intArr[i]
		} else if intArr[i] < sn {
			sn = intArr[i]
		}
	}
	if fn*sn > fp*sp {
		fmt.Print(fn, sn)
	} else {
		fmt.Print(sp, fp)
	}
}
