package main

import (
	"fmt"

	"github.com/codegangsta/inject"
)

func test() {
	fmt.Println("123")
}

func add(num1 int32, num2 int32) int32 {
	fmt.Println(num1, num2)
	return num1 + num2
}

func main() {
	injector := inject.New()
	injector.Invoke(test)

	var a int32 = 10

	injector.Map(a)

	rlt, err := injector.Invoke(add)

	if err == nil {
		fmt.Println(rlt[0])
	} else {
		fmt.Println(err)
	}
}
