# 闭包实现工厂模式

## 场景

在游戏中创建两个角色，战士和法师，如果不使用闭包或者工厂模式，我们可以需要写如下的代码

```go
package main

import "fmt"

func main() {
	name, carrer, gender, hp, mp := createPlayer("西瓜太郎", "战士", "男")
	fmt.Println(name, carrer, gender, hp, mp)
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
```



## 工厂模式

上述代码中这种通过传参数不同自动生成数据对象的设计模式就是工程模式，不过我们可以通过**闭包**对其进行进一步的简化



## 闭包

从上面的需求来看，其实相同职业的`hp`和`mp`是相同的，唯一不同的只有名称和职业，所以可以创建一个只创建战士或者法师的函数进行调用，如下

```go
package main

import "fmt"

func main() {
	var warriorCtr = playerCreate("战士")
	var masterCtr = playerCreate("法师")

	fmt.Println(warriorCtr("西瓜太郎2", "男"))
	fmt.Println(warriorCtr("西瓜太郎3", "女"))
	fmt.Println(masterCtr("西瓜太郎4", "女"))
}

func playerCreate(career string) func(name string, gender string) (string, string, string, int32, int32) {
	var hp int32 = 0
	var mp int32 = 0

	switch career {
	case "战士":
		hp = 150
		mp = 80
	case "法师":
		hp = 80
		mp = 200
	}

	return func(name string, gender string) (string, string, string, int32, int32) {
		return name, career, gender, hp, mp
	}
}
```

从上述可以见，当我们需要创建男战士、女战士的时候，只需要将对应的名称和性别传入即可。

上述代码中，`hp`和`mp`都是用的`playerCreate`上下文中的参数，这种方式就称为闭包。

**闭包的优点**：

1. 每个变量都互相独立，且有记忆效应，存储各自的状态互不影响，保护闭包数据不受影响

2. 由于返回的值可以赋值给外部，这样就可以方便**后期更改函数中变量的值**，相较于普通函数的执行完毕后无法更改，具有记忆性
