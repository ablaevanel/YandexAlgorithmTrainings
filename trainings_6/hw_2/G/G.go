package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Buffer(make([]byte, 0, 1024*1024*10), 1024*1024*10)
	scanner.Scan()
	nums := strings.Split(strings.Trim(scanner.Text(), " "), " ")
	n, _ := strconv.ParseUint(nums[0], 10, 64)
	c, _ := strconv.ParseUint(nums[1], 10, 64)
	scanner.Scan()
	s := strings.Trim(scanner.Text(), " ")
	m := map[byte]uint64{
		'a': 0,
		'b': 0,
	}
	var i, j uint64 = 0, 0
	var length, totalB, lastA, currC uint64 = 0, 0, 0, 0
	for i < n && j < n && i <= j {
		if s[j] == 'a' {
			if m['b'] != 0 {
				currC += m['a'] * m['b']
				totalB += m['b']
				m['b'] = 0
				lastA = j
			}
			m['a']++
		} else if s[j] == 'b' {
			m['b']++
		}
		for m['a']*m['b']+currC > c && i <= j {
			if s[i] == 'a' {
				m['a']--
				currC -= totalB
			} else if s[i] == 'b' {
				if i < lastA {
					totalB--
				} else {
					m['b']--
				}
			}
			i++
		}
		if j-i+1 > length {
			length = j - i + 1
		}
		j++
	}
	output, _ := os.Create("output.txt")
	defer output.Close()
	output.WriteString(strconv.Itoa(int(length)))
}
