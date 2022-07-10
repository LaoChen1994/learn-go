package main

import (
	"fmt"
	"reflect"
)

func main() {
	var exampleVal int = 10

	exampleValAddr := &exampleVal

	exampleValReal := *&exampleVal

	// 0xc000012098
	fmt.Println(exampleValAddr)
	// 10
	fmt.Println(exampleValReal)

	newExampleArr := new(int64)
	*newExampleArr = 100

	fmt.Println(*newExampleArr)

	num1 := 10
	num2 := 3

	fmt.Println(num1 / num2) // 3
	fmt.Println(num1 % num2) // 1
	fmt.Println(num1 == num2)

	// 1010
	// 0011
	fmt.Println(num1 & num2) // 2
	fmt.Println(num1 | num2) // 11

	bool1 := true
	bool2 := false

	fmt.Println(bool1 && bool2)
	fmt.Println(bool1 || bool2)

	val1 := 3.22
	fmt.Println(reflect.TypeOf(val1))

	val2 := int32(val1)

	fmt.Println(reflect.TypeOf(val2), val2)
}
