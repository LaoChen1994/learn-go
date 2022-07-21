# 结构体二—构建有生命的对象

在`go`语言中使用结构体来实现构造函数和方法，**go语言中本来是没有构造函数的，但是可以通过结构体初始化的过程来实现构造函数**


## 面向对象语言中的构造函数

**构造函数申明方法**：（1） 申明结构体 （2）申明一个函数，其返回值是生成结构体实例的引用 （3）在构造函数中可以通过转换来对传入结构体中的值进行特殊的封装

**注意点**：之前引用传值的时候学过，传入或者传出实参前增加`&`，表示其传递的参数是引用，在形参前添加`*`代表该形参是传递的是引用类型

**例子**

```go
package main

import "fmt"

type Dog struct {
	Breed  string
	Age    int32
	Weight float64
	Gender string
}

// 这里申明类型的时候代表扔出去的是个引用类型
func NewDog(Breed string, Age int32, Weight float64, Gender string) *Dog {
	var gender int32 = 0

    // 对参数进行特殊的封装
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

func main() {
	dog1 := NewDog("Haba1", 12, 10.2, "male")
	dog2 := NewDog("Haba2", 13, 10.2, "female")

	// &{Haba1 12 10.2 2}
	fmt.Println(dog1)
	// &{Haba2 13 10.2 1}
	fmt.Println(dog2)
}
```

## 面向对象语言中的方法

### 接收器的定义和使用

**作用**：对于一个对象而言，除了要有静态属性，还需要有**方法**，来为一个对象关联上对应的方法

**定义方法**：`func (struct params type) func_name (props) (return_type) {}`

**注意**：相较于普通的函数，接收器函数新增了一个接收器变量类型的定义方式

**例子**

```go
type Dog struct {
	Breed  string
	Age    int32
	Weight float64
	Gender int32
}

// 会将这个方法直接绑到Dog这个结构体上去
func (d *Dog) GetGender() string {
	if d.Gender == 1 {
		return "female"
	} else {
		return "male"
	}
}

func main() {
	// 这个参考第一节中
	dog1 := NewDog("Haba1", 12, 10.2, "male")
	dog2 := NewDog("Haba2", 13, 10.2, "female")

	fmt.Println(dog1)
	fmt.Println(dog2)

	fmt.Println(dog1.GetGender())
}
```

### 接收器中的引用传值和值传递

**区别**：使用引用传值会直接影响到原有结构体中定义的值，而值传递因为是生成副本所以就不会影响

**例子**

由于growUp2中是值类型，所以其改变不会改动到实际`dog2`的`age`变量

```go
func (d *Dog) growUp1() {
	d.Age++
}

func (d Dog) growUp2() {
	d.Age++
}

func main()  {
	// 省略上述相同代码
	// 12
	fmt.Println(dog1.Age)

	dog1.growUp1()

	// 13
	fmt.Println(dog1.Age)

	dog1.growUp2()

	// 13
	fmt.Println(dog1.Age)
}
```