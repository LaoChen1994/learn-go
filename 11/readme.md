## 结构体一：万物皆可结构体

## GO语言中结构体的定义和使用

**什么是结构体**：结构体是由一系列具有相同类型或不同类型的数据结构构成的数据集合

**定义方式**：`type StructName struct {}`

**初始化**：通过`struct_name { key: val }`来进行初始化，如果需要初始化结构体内所有参数，可以使用`struct_name { value, value2, value3 }`

**修改和使用参数**：使用`struct_name.prop = val`的方式进行修改，通过`struct_name.prop`来实现值的获取

**例子**

```go
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

    // 也可以在初始化的时候对一部分结构体变量进行赋值
    // 这个时候需要带上参数名
    player2 := Player{
		name:   "pidan2",
		career: "master",
	}

	player2.hp = 80
	player2.mp = 300

	fmt.Println(player2)

}
```

## 匿名结构体的定义和使用

使用结构体的目的是：将事物的属性进行**抽象**，将复杂的问题简单化，便于代码的**扩展性**以及**复用性**

但是如果，某一个结构体只在非常小的范围内使用，且不具有复用性，**可以通过匿名结构体的方式来进行赋值**

```go
bike := struct {
		brand string
		price float32
	}{
		brand: "凤凰牌",
	}

bike.price = 399.0
// {凤凰牌 399}
fmt.Println(bike)
```



匿名结构体是序列化和反序列化的基础，在后面会着重进行讨论
