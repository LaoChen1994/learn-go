# 反射（三）控制反转与依赖注入

## 概念

### 控制反转

普通函数调用方式：调用方直接调用某个函数例如`test()`

**控制反转**：发起主动调用的一方不再直接调用某个函数，通过某些配置，将配置信息交给框架，完成函数调用

**优势**：实现调用方和执行方的解耦，让程序更加自由化、多样化

### 依赖注入

什么是依赖注入：依赖注入是实现控制反转的一种方法，其目的是让调用方将方法动态注入到执行方，为执行方动态执行做铺垫

使用`github.com/codegangsta/inject`这个包来实现依赖的注入

1. 引入`github.com/codegangsta/inject`
   ```bash
    go get github.com/codegangsta/inject
   ```
2. 创建`Injector`，使用`inject.New()`
3. 使用`Invoke`方法来让执行器执行某个函数

```go
package main
import (
	"fmt"

	"github.com/codegangsta/inject"
)

func test() {
	fmt.Println("123")
}

func add(num1 int32, num2 int32) int32 {
	return num1 + num2
}


func main() {
	injector := inject.New()
	injector.Invoke(test)

    // 有参数注入的方法
    var a int32 = 10

	injector.Map(a)

	rlt, err := injector.Invoke(add)

	if err == nil {
		fmt.Println(rlt[0])
	} else {
		fmt.Println(err)
	}
}
```

## 总结

上述的反转控制只是一个例子，因为`inject`这个包我感觉也不是特别完善，引入其通过`Map`做依赖参数的注入是通过`Type`类型，所以如果`num1`和`num2`是相同类型的话，可能会导致注入相同的值，这里只是给了一个反转控制大致例子的思想

思想是：**因为需要做依赖的注入，所以无法在编译时确定，一定是运行时**，所以这个时候需要用到反射 + interface 泛型

