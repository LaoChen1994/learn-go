# 指针和运算符

## 

## 指针类型

GO的指针：**类型指针**和**切片指针**

**类型指针**：传递数据时使用指针可以避免创建数据副本，节约内存开销，该指针不能偏移取数（不然会非法修改其他数据），更有利于垃圾回收

**切片指针**：由**指向起始元素指针**、**元素数量**和**总容量构成**，如果切片发生越界，会发生宕机。



获取指针的方式：

1. 使用`&`来获取变量地址

2. 用`*`来获取某个变量地址对应的值，所以`*&`一般连用，用于拿到对应地址里面存的值



## 类型指针取数例子

```go
package main

import "fmt"

func main() {
	var exampleVal int = 10

	exampleValAddr := &exampleVal

	exampleValReal := *&exampleVal

	// 0xc000012098
	fmt.Println(exampleValAddr)
	// 10
	fmt.Println(exampleValReal)
}

```



## 使用new来创建指针变量
