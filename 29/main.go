package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int
	Gender int
}

type myint int

func main() {
	personExample := &Person{
		Name:   "小王",
		Age:    18,
		Gender: 1,
	}

	personType := reflect.TypeOf(personExample)

	fmt.Println(personType.Name())
	fmt.Println(personType.Kind())

	fmt.Println(personType.Elem().NumField())

	values := reflect.ValueOf(personExample).Elem()
	personElem := personType.Elem()

	for i := 0; i < values.NumField(); i++ {
		field := values.Field(i)
		name := personElem.Field(i).Name
		t := personElem.Field(i).Type

		fmt.Println(name, t, field)
	}

	fmt.Println("age =>", values.FieldByName("Age"))
	values.FieldByName("Name").SetString("小红")
	values.FieldByName("Age").SetInt(20)
	fmt.Println(values.FieldByName("Name").IsValid(), values.FieldByName("Name").CanSet())
	fmt.Println("name =>", values.FieldByName("Name"))
	fmt.Println("age =>", values.FieldByName("Age"))

	// 调用函数
	accFuncType := reflect.TypeOf(addCalc)
	fmt.Println(accFuncType.Kind())

	accFuncParam := []reflect.Value{reflect.ValueOf(20), reflect.ValueOf(200)}

	accValue := reflect.ValueOf(addCalc)

	res := accValue.Call(accFuncParam)

	fmt.Println(res[0])

	// 创建实例
	var num myint = 100

	typeOfNum := reflect.TypeOf(num)
	another := reflect.New(typeOfNum)

	another.Elem().SetInt(200)

	fmt.Println(another.Elem().Int())
	fmt.Println(num)
	fmt.Println(another.Type().Kind(), typeOfNum.Kind())
}

func addCalc(num1 int, num2 int) int {
	return num1 + num2
}
