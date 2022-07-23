# 接口（一）——结构的定义和使用

## 为什么要用结构

结构就是规定了某个对象的行为，接口的使用者（调用方）无需关注某个动作的具体实现；实现者（提供方）无需关心外部如何使用该接口

接口的关键步骤：
    1. 接口的定义
    2. 接口的实现


## 接口的定义

**定义方式**：使用`type interface_name interface {}`来进行定义

**注意**：通过`interface`申明接口时候，可以申明一个或多个方法。

**例子**

```go
type LearnCode interface {
	LearnGo() string
    StartCoding(code string) string
    FindWork(company string) string
}

```

## 接口的实现

**实现方式**：`func (interface_name interface type) function_name(params)[return_values]{}`

**注意**：
    1. 接口定义的方法和实现接口的类型方法格式一致,包括`方法名`、`函数名`、`返回值`
    2. 接口中所有定义的方法都需要实现(不然在接口赋值的时候会报错)
    3. 如果是接口除了方法以外，还规定了相应的内部参数，这里就需要注意了，可以通过一个`init`方法来对基础值进行初始化

```go
package main

import "fmt"

type LearnCode interface {
	LearnBasic() string
	StartCoding(code string) string
	FindWork(company string) string
	init()
}

type StudyGo struct {
	language string
}

func (coder StudyGo) LearnBasic() string {
	return "学习" + coder.language
}

func (coder StudyGo) StartCoding(code string) string {
	return "使用" + coder.language + "写" + code
}

func (coder StudyGo) FindWork(company string) string {
	return "去" + company + "写" + coder.language
}

func (coder *StudyGo) init() {
	coder.language = "go"
}
```

## 接口的调用

**使用方法**：1. 申明一个变量类型为`interface type`的变量。2. 使用`new(struct_type)`来实例化对象，之前说过通过`new(type)`可以在内存中为变量申请一个存放地址，这里就是申请新的`struct_type`的地址

**注意**：一种类型可以实现多个接口，多种类型可以实现相同的接口

```go
func main() {
	var goer LearnCode

	goer = new(StudyGo)
	goer.init()

	fmt.Println(goer.LearnBasic())
	fmt.Println(goer.StartCoding("Hello World"))
	fmt.Println(goer.FindWork("bytedance"))
}
```