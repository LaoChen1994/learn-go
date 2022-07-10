# 基础语法：变量、常量、声明和基本类型

## 变量和常量

变量：后面可以修改的

常量：声明后无法修改的



## 推荐的命名规范

变量：驼峰命名，如果是布尔类型建议以`can`、`allow`、`has`、`can`来开头

常量：应用全用大写字母来命名，单词间用`_`做分割



## 声明与赋值

命名格式：`var name type`，`const name type = value`

例子：`var num int = 100`,



命名简写的方式：`str := "你好啊"`，对于们没有命名的变量，可以通过这种方式来进行命名。



## 变量的批量声明和赋值

```go
package constant

import "fmt"

func main() {
	var (
		age    int    = 18
		name   string = "xx"
		salary int    = 1000
	)

	fmt.Print("name ->", name, "age ->", age, "salary ->", salary)
}


```



## 基本数据类型

GO的数据类型：基本数据类型、接口类型和复合类型（数组、切片、结构体）。

基本数据类型：

+ int类型：int8，int16等

+ 浮点数类型：float32、float64等

+ 复数类型：complex

+ 布尔类型：bool

+ 字符串类型：string



需要注意的是在字符串类型中，可以使用`\`做字符的转义，另外可以使用`\n`来做换行
