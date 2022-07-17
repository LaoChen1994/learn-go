# 匿名函数

## 定义和使用

**作用**：匿名函数是实现go中回调的核心技能，通过回调来保证程序运行的正确性和及时性

**定义方式**：`func ([params_list])([return_value_list]) {}`

**使用方式**

1. 使用IIFE的方式进行调用
   
   ```go
   package main
   
   import "fmt"
   
   func main() {
   	func(text string) {
   		fmt.Println(text)
   	}("立即使用")
   }
   
   ```

2. 赋值给变量调用
   
   ```go
   stringLoop := func(text string) {
   		fmt.Println(text)
   	}
       
   stringLoop("传参使用")
   ```



## 回调使用实战

### 模拟下载成功的回调函数

1. 回调函数在参数中的声明，用`func()`来标记他是一个函数

2. 使用全局变量`precent`来计算进度，使用闭包的方式来实现这样的功能

```go
package main

import (
	"fmt"
	"time"
)

var precent int = 0

func main() {
    // 使用go关键词来开启多线程
	go download("", func() {
		fmt.Println("下载成功")
		isEnd = true
	})

    // 这里需要使用轮询的方式不断查询，不然的话进程会直接结束，就不会有下载成功了
	for {
		if isEnd {
			break
		} else {
			time.Sleep(5 * time.Second)
			fmt.Println("当前进度", precent, "%")
		}
	}

}

func download(url string, onSuccess func()) {
	for {
		time.Sleep(1 * time.Second)
		precent += 10

		if precent >= 100 {
			onSuccess()
			break
		}
	}
}


```
