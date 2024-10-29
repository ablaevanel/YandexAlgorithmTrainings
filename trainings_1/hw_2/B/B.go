package main

import "fmt"

const CONSTANT = -2000000000

func main() {
	var curr, prev, cnt int
	fmt.Scan(&prev, &curr)
	cnt = 1
	t := map[string]bool{
		"CONSTANT":          true,
		"ASCENDING":         true,
		"WEAKLY ASCENDING":  true,
		"DESCENDING":        true,
		"WEAKLY DESCENDING": true,
	}
	for curr != CONSTANT {
		if prev < curr {
			t["CONSTANT"] = false
			t["DESCENDING"] = false
			t["WEAKLY DESCENDING"] = false
		} else if prev == curr {
			t["DESCENDING"] = false
			t["ASCENDING"] = false
		} else {
			t["CONSTANT"] = false
			t["ASCENDING"] = false
			t["WEAKLY ASCENDING"] = false
		}
		prev = curr
		fmt.Scan(&curr)
		cnt++
	}
	if t["CONSTANT"] {
		fmt.Print("CONSTANT")
		return
	}
	if t["DESCENDING"] && t["WEAKLY DESCENDING"] {
		fmt.Print("DESCENDING")
		return
	}
	if t["ASCENDING"] && t["WEAKLY ASCENDING"] {
		fmt.Print("ASCENDING")
		return
	}
	ans := ""
	for k, v := range t {
		if v {
			ans = k
		}
	}
	if ans == "" {
		ans = "RANDOM"
	}
	fmt.Print(ans)
}
