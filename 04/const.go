package main

import "fmt"

const TEMP = 37

func main() {
	var num int16 = 100
	str := "你好啊"
	fmt.Print(num)
	fmt.Printf(str)

	var (
		age    int       = 18
		name   string    = "xx"
		salary int       = 1000
		comp   complex64 = 1 + 2i
	)

	fmt.Print("name ->", name, "age ->", age, "salary ->", salary, comp)
	fmt.Print("\ntoday temp is ", TEMP)
}
