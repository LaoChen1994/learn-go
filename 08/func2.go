package main

import (
	"fmt"
	"reflect"
)

func main() {
	x2, x3 := getX2X3(4)

	var rlt = 10

	reply(1, 2, &rlt)

	fmt.Println(x2, x3)

	fmt.Println(rlt)
	fmt.Println(min(8, 2, 9, 4, 5, 1))

	var a int32 = 1
	typecheck(a, "123", []int{1, 2, 3}, true)
}

func getX2X3(num int) (x2 int, x3 int) {
	x2 = num * 2
	x3 = num * 3
	return
}

func mult_returnval(a int32, b int32) (sum int32, diff int32) {
	sum = a + b
	diff = a - b
	return
}

func reply(a, b int, rlt *int) {
	*rlt = a * b
}

func min(a ...int) (val int) {
	val = 0

	if len(a) == 0 {
		return
	}

	val = a[0]

	for _, v := range a {
		if v < val {
			val = v
		}
	}

	return
}

func typecheck(props ...interface{}) {
	for _, value := range props {
		switch reflect.ValueOf(value).Kind() {
		case reflect.String:
			fmt.Println(value, "is", " string")
		case reflect.Int32:
			fmt.Println(value, "is", " int32")
		case reflect.Array:
			fmt.Println(value, "is", " array")
		case reflect.Bool:
			fmt.Println(value, "is", " boolean")
		default:
			fmt.Println(value, " is", reflect.ValueOf(value).Kind())
		}
	}
}
