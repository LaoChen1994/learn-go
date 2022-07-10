package main

import (
	"fmt"
)

func main() {
	// trigle()
	findNum()

}

func findNum() {
	flag := false

	for num := 2; num < 100; num++ {
		if num == 2 {
			fmt.Println(num)
			continue
		} else {
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = true
					break
				}
			}
		}

		if !flag {
			fmt.Println(num)
		}

		flag = false
	}
}

func trigle() {
	n := 7
	mode := "down"

	if mode == "all" || mode == "up" {
		for i := 1; i <= n; i++ {
			for j := 0; j < n-i; j++ {
				fmt.Print(" ")
			}

			for k := 0; k < 2*i-1; k++ {
				fmt.Print("*")
			}

			fmt.Println("")
		}
	}

	if mode == "down" || mode == "all" {
		for i := n - 1; i > 0; i-- {
			for j := 0; j < n-i; j++ {
				fmt.Print(" ")
			}

			for k := 0; k < 2*i-1; k++ {
				fmt.Print("*")
			}

			fmt.Println("")
		}
	}
}
