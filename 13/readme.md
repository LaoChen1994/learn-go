# 结构体三——继承

## 继承的实现与使用

**什么是继承**：通过继承，子类能在父类的基础上扩展新的功能，且对抽象的父类功能进行复用

**如何实现**：go语言中的继承是通过结构体的嵌套来实现的

**例子**

```go
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

func main() {
	bird := NewBird("小鸟", 3, "male", "blue")

	fmt.Println(bird)
	bird.CommonAnimal.Eat()
	bird.Fly()
}

```

从上面例子中，如果我们同样创建一个狗的实体，继承于`Animal`，其确实`Animal`上的`Eat`方法是能共享的，但是`Fly`方法因为是绑定在`Bird`上的，所以并不能对`Dog`生效


## 匿名结构体嵌套

```go
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
```

使用这种匿名的方式对结构体进行声明，然后通过调用`dog.string`的方式来获取小狗毛发的颜色，但是这样子来看无法知道`dog.string`到底对应什么变量, 且有多个`string`变量的话就无法进行重复申明