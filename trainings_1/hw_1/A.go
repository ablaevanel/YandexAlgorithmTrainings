package main

import "fmt"

func main() {
	var troom, tcond int
	var mode string
	fmt.Scanln(&troom, &tcond)
	fmt.Scanln(&mode)
	switch mode {
	case "heat":
		if troom < tcond {
			fmt.Print(tcond)
		} else {
			fmt.Print(troom)
		}
	case "freeze":
		if troom > tcond {
			fmt.Print(tcond)
		} else {
			fmt.Print(troom)
		}
	case "auto":
		fmt.Print(tcond)
	case "fan":
		fmt.Print(troom)
	}
}
