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
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 1024 * 1024 * 1024
	scanner.Buffer(make([]byte, 0, maxCapacity), maxCapacity)
	scanner.Scan()
	arr := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	intArr := make([]int, len(arr))
	for i, str := range arr {
		intArr[i], _ = strconv.Atoi(str)
	}
	sort.Ints(intArr[:3])
	tp, sp, fp := intArr[0], intArr[1], intArr[2]
	sn, fn := sp, tp
	for i := 3; i < len(intArr); i++ {
		if intArr[i] > fp {
			tp = sp
			sp = fp
			fp = intArr[i]
		} else if intArr[i] > sp {
			tp = sp
			sp = intArr[i]
		} else if intArr[i] > tp {
			tp = intArr[i]
		}
		if intArr[i] < fn {
			sn = fn
			fn = intArr[i]
		} else if intArr[i] < sn {
			sn = intArr[i]
		}
	}
	if fn*sn*fp > fp*sp*tp {
		fmt.Print(fn, sn, fp)
	} else {
		fmt.Print(tp, sp, fp)
	}
}
