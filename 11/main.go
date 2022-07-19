package main

import "fmt"

type Player struct {
	name   string
	career string
	hp     int32
	mp     int32
}

func main() {
	player := Player{
		"pidan",
		"warrior",
		200,
		80,
	}

	// {pidan warrior 200 80}
	fmt.Println(player)

	player.career = "king of warrior"
	player.hp += 50
	player.mp += 30

	// {pidan king of warrior 250 110}
	fmt.Println(player)

	player2 := Player{
		name:   "pidan2",
		career: "master",
	}

	player2.hp = 80
	player2.mp = 300

	fmt.Println(player2)

	bike := struct {
		brand string
		price float32
	}{
		brand: "凤凰牌",
	}

	bike.price = 399.0
	fmt.Println(bike)
}
