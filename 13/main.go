package main

import "fmt"

type Animal struct {
	Name   string
	Age    int
	Gender string
}

func (animal Animal) Eat() {
	fmt.Println(animal.Name, "正在觅食...")
}

type Bird struct {
	WingColor    string
	CommonAnimal Animal
}

func NewBird(name string, age int, gender string, wingColor string) *Bird {
	return &Bird{
		wingColor,
		Animal{
			Name:   name,
			Age:    age,
			Gender: gender,
		},
	}
}

func (b *Bird) Fly() {
	fmt.Println(b.CommonAnimal.Name, "开始飞啦~")
}

type Dog struct {
	string
	Animal
}

func (d *Dog) Run() {
	fmt.Println(d.Name, "正在奔跑")
}

func NewDog(breed string, name string, age int, gender string) *Dog {
	return &Dog{
		breed,
		Animal{
			name,
			age,
			gender,
		},
	}
}

func main() {
	bird := *NewBird("小鸟", 3, "male", "blue")

	fmt.Println(bird)
	bird.CommonAnimal.Eat()
	bird.Fly()

	dog := *NewDog("black", "小狗", 3, "female")

	dog.string = "white"

	fmt.Println(dog)
}
