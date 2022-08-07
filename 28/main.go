package main

import (
	"fmt"
	"reflect"
)

func main() {
	var testNum int = 10
	testNumPtr := &testNum
	typeOfNum := reflect.TypeOf(testNumPtr)
	fmt.Println(typeOfNum.Name())
	fmt.Println(typeOfNum.Kind())

	num := reflect.ValueOf(testNum)

	fmt.Println(num.CanSet())
	if num.CanSet() {
		num.SetInt(30)
	}

	fmt.Println(testNum)

	value := reflect.ValueOf(testNumPtr)

	fmt.Println(value.Elem().CanSet(), value.Elem().CanAddr())

	fmt.Println(value.IsValid())
	fmt.Println(value.IsNil())

	if value.IsNil() || !value.IsValid() {
		return
	}

	fmt.Println(value.Elem().Int(), testNum, *testNumPtr)

	value.Elem().SetInt(20)

	fmt.Println(value.Elem().Int(), testNum, *testNumPtr)
	// val := value.Int()

	// fmt.Println(val)
}
