package main

import "fmt"

type Dog struct {
	Breed  string
	Age    int32
	Weight float64
	Gender int32
}

// 这里申明类型的时候代表扔出去的是个引用类型
func NewDog(Breed string, Age int32, Weight float64, Gender string) *Dog {
	var gender int32 = 0

	if Gender == "female" || Gender == "母" {
		gender = 1
	} else if Gender == "male" || Gender == "公" {
		gender = 2
	}

	// 这里代表传出去的是数据的引用
	return &Dog{
		Breed:  Breed,
		Age:    Age,
		Weight: Weight,
		Gender: gender,
	}
}

func (d Dog) GetGender() string {
	if d.Gender == 1 {
		return "female"
	} else {
		return "male"
	}
}

func (d *Dog) growUp1() {
	d.Age++
}

func (d Dog) growUp2() {
	d.Age++
}

func main() {
	dog1 := NewDog("Haba1", 12, 10.2, "male")
	dog2 := NewDog("Haba2", 13, 10.2, "female")

	fmt.Println(dog1)
	fmt.Println(dog2)

	fmt.Println(dog1.GetGender())

	fmt.Println(dog1.Age)

	dog1.growUp1()

	fmt.Println(dog1.Age)

	dog1.growUp2()

	fmt.Println(dog1.Age)
}
