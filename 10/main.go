package main

import "fmt"

func main() {
	name, carrer, gender, hp, mp := createPlayer("西瓜太郎", "战士", "男")
	fmt.Println(name, carrer, gender, hp, mp)

	var warriorCtr = playerCreate("战士")
	var masterCtr = playerCreate("法师")

	_, _, _, h1, m1 := warriorCtr("西瓜太郎2", "男")
	_, _, _, h2, m2 := warriorCtr("西瓜太郎3", "女")

	fmt.Println("h1 ->", h1, "h2 ->", h2)
	fmt.Println("m1 ->", m1, "m2 ->", m2)

	h1 -= 20
	m2 -= 15

	fmt.Println("h1 ->", h1, "h2 ->", h2)
	fmt.Println("m1 ->", m1, "m2 ->", m2)

	fmt.Println(masterCtr("西瓜太郎4", "女"))
}

func createPlayer(name string, career string, gender string) (string, string, string, int32, int32) {
	var hp int32 = 0
	var mp int32 = 0

	switch career {
	case "战士":
		hp = 150
		mp = 80
	case "法师":
		hp = 100
		mp = 100
	}

	return name, career, gender, hp, mp
}

func playerCreate(career string) func(name string, gender string) (string, string, string, int32, int32) {
	var hp int32 = 0
	var mp int32 = 0

	switch career {
	case "战士":
		hp = 150
		mp = 80
	case "法师":
		hp = 100
		mp = 100
	}

	return func(name string, gender string) (string, string, string, int32, int32) {
		return name, career, gender, hp, mp
	}
}
