package main

import (
	"fmt"
	"strings"
)

func strFunc(val []string) string {
	return strings.Join(val, ",")
}

func main() {
	defer fmt.Println("我是最后一个出来的")
	defer fmt.Println("我是倒数第二个出来的")
	defer fmt.Println("我是倒数第三个出来的")

	var list []string
	list = append(list, "123")
	list = append(list, "456")
	res := strFunc(list)

	var rlt []int

	findPrimeNumber(&rlt, 100)

	fmt.Println(res)
	fmt.Println(rlt)

	val := calcFactorial(5)

	var c int

	c = calcFibonacci(10)

	b := *&c

	fmt.Println(val)
	fmt.Println(b)
}

func findPrimeNumber(slice *[]int, max int) {
	for i := 1; i < max; i++ {
		if i == 2 {
			*slice = append(*slice, i)
			continue
		}

		flag := false

		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = true
				break
			}
		}

		if !flag {
			*slice = append(*slice, i)
		}

		flag = false
	}
}

func calcFactorial(n int) (result int) {
	if n > 0 {
		return n * calcFactorial(n-1)
	}

	return 1
}

func calcFibonacci(n int) (result int) {
	if n == 1 || n == 2 {
		return 1
	}

	return calcFibonacci(n-1) + calcFibonacci(n-2)
}
