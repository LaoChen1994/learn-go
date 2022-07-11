package main

import "fmt"

func main() {
	var array [4]int
	var slice []int

	arrayIndex := 0
	flag := false

	for i := 2; i <= 10; i++ {
		if i == 2 {
			array[arrayIndex] = i
			slice = append(slice, i)
			arrayIndex++
			continue
		}
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = true
				break
			}
		}

		if !flag {
			array[arrayIndex] = i
			slice = append(slice, i)
			arrayIndex++
		}

		flag = false
	}

	fmt.Println(array)
	fmt.Println(slice)

	mapExample()
}

func mapExample() {
	var stuInfo = make(map[string]string)

	stuInfo["001"] = "zhangSan"
	stuInfo["002"] = "liSi"
	stuInfo["003"] = "wangWu"

	fmt.Println(stuInfo)
}
