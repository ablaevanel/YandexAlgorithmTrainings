package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	letter := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&letter[i])
	}
	cuttedLetter := make([]string, 0, n)
	cuttedLetter = append(cuttedLetter, letter[0])
	for i := 1; i < n; i++ {
		if letter[i] != cuttedLetter[len(cuttedLetter)-1] {
			cuttedLetter = append(cuttedLetter, letter[i])
		}
	}
	flag := true
	for i := 0; i < n; i++ {
		if string(cuttedLetter[0][i]) == "#" {
			flag = false
		}
	}
	if flag {
		if len(cuttedLetter) == 1 {
			fmt.Print("X")
			return
		}
		cuttedLetter = cuttedLetter[1:]
	}
	flag = true
	for i := 0; i < n; i++ {
		if string(cuttedLetter[len(cuttedLetter)-1][i]) == "#" {
			flag = false
		}
	}
	if flag {
		if len(cuttedLetter) == 1 {
			fmt.Print("X")
			return
		}
		cuttedLetter = cuttedLetter[:len(cuttedLetter)-1]
	}
	reversedLetter := make([][]string, n)
	for i := 0; i < n; i++ {
		reversedLetter[i] = make([]string, len(cuttedLetter))
	}
	for i := 0; i < len(cuttedLetter); i++ {
		for j := 0; j < n; j++ {
			reversedLetter[j][i] = string(cuttedLetter[i][j])
		}
	}
	reversedLetter2 := make([]string, n)
	for i := 0; i < n; i++ {
		row := ""
		for j := 0; j < len(reversedLetter[0]); j++ {
			row += reversedLetter[i][j]
		}
		reversedLetter2[i] = row
	}
	cuttedLetter2 := make([]string, 0, n)
	cuttedLetter2 = append(cuttedLetter2, reversedLetter2[0])
	for i := 1; i < n; i++ {
		if reversedLetter2[i] != cuttedLetter2[len(cuttedLetter2)-1] {
			cuttedLetter2 = append(cuttedLetter2, reversedLetter2[i])
		}
	}
	flag = true
	for i := 0; i < len(cuttedLetter2[0]); i++ {
		if string(cuttedLetter2[0][i]) == "#" {
			flag = false
		}
	}
	if flag {
		if len(cuttedLetter2) == 1 {
			fmt.Print("X")
			return
		}
		cuttedLetter2 = cuttedLetter2[1:]
	}
	flag = true
	for i := 0; i < len(cuttedLetter2[0]); i++ {
		if string(cuttedLetter2[len(cuttedLetter2)-1][i]) == "#" {
			flag = false
		}
	}
	if flag {
		if len(cuttedLetter2) == 1 {
			fmt.Print("X")
			return
		}
		cuttedLetter2 = cuttedLetter2[:len(cuttedLetter2)-1]
	}
	reversedLetter3 := make([][]string, len(cuttedLetter2[0]))
	for i := 0; i < len(reversedLetter3); i++ {
		reversedLetter3[i] = make([]string, len(cuttedLetter2))
	}
	for i := 0; i < len(cuttedLetter2); i++ {
		for j := 0; j < len(cuttedLetter2[0]); j++ {
			reversedLetter3[j][i] = string(cuttedLetter2[i][j])
		}
	}
	reversedLetter4 := make([]string, len(reversedLetter3))
	for i := 0; i < len(reversedLetter3); i++ {
		row := ""
		for j := 0; j < len(reversedLetter3[0]); j++ {
			row += reversedLetter3[i][j]
		}
		reversedLetter4[i] = row
	}
	if len(reversedLetter4) == 1 && reversedLetter4[0] == "#" {
		fmt.Print("I")
	} else if len(reversedLetter4) == 2 && reversedLetter4[0] == "#." && reversedLetter4[1] == "##" {
		fmt.Print("L")
	} else if len(reversedLetter4) == 3 {
		if len(reversedLetter4[0]) == 2 {
			if reversedLetter4[0] == "##" && reversedLetter4[1] == "#." && reversedLetter4[2] == "##" {
				fmt.Print("C")
			} else {
				fmt.Print("X")
			}
		} else if len(reversedLetter4[0]) == 3 {
			if reversedLetter4[0] == "###" && reversedLetter4[1] == "#.#" && reversedLetter4[2] == "###" {
				fmt.Print("O")
			} else if reversedLetter4[0] == "#.#" && reversedLetter4[1] == "###" && reversedLetter4[2] == "#.#" {
				fmt.Print("H")
			} else {
				fmt.Print("X")
			}
		} else {
			fmt.Print("X")
		}
	} else if len(reversedLetter4) == 4 {
		if reversedLetter4[0] == "###" && reversedLetter4[1] == "#.#" && reversedLetter4[2] == "###" && reversedLetter4[3] == "#.." {
			fmt.Print("P")
		} else {
			fmt.Print("X")
		}
	} else {
		fmt.Print("X")
	}
}
