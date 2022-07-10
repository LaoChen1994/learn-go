package main

import "fmt"

func main() {
	var exampleVal int = 10

	exampleValAddr := &exampleVal

	exampleValReal := *&exampleVal

	// 0xc000012098
	fmt.Println(exampleValAddr)
	// 10
	fmt.Println(exampleValReal)
}
