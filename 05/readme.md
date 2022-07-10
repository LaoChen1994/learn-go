# 指针和运算符

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

## 使用new来创建指针变量地址

方法：通过`new(type)`的方式来在内存空间中为变量**创建对应的地址**，这个时候变量是没有被赋值的，需要通过`*`来为对应的来对对应的变量进行赋值



## 运算符与优先级

1. 算数运算符，加减乘除和数学上的优先级一样

2. 除法`int`类型只保留整数，不会有小数

3. 取余需要使用`%`来进行

```go
	num1 := 10
	num2 := 3

	fmt.Println(num1 / num2) // 3
	fmt.Println(num1 % num2) // 1
```



## 关系运算符

关系运算符包括：`>`，`<`，`>=`，`==`



## 逻辑运算符

`&&`，`||`，`!`

```go
	bool1 := true
	bool2 := false

	fmt.Println(bool1 && bool2)
	fmt.Println(bool1 || bool2)
```



## 位运算符

使用位运算来做算法的加减

```go
	num1 := 10
	num2 := 3

	// 1010
	// 0011
	fmt.Println(num1 & num2) // 2
	fmt.Println(num1 | num2) // 11
```



## 类型转换

转换的方式：`val2 = type(val1)`

1. 只能将相同数据类型的进行转换，比如`int32`转换成`int64`，如果将`bool`转换成`string`会产生报错

2. 如果将取值范围较大的转换成取值范围较小的类型会发生精度丢失的情况，比如`float32`转换成`int32`

```go
val1 := 3.22
fmt.Println(reflect.TypeOf(val1))
val2 := int32(val1)
fmt.Println(reflect.TypeOf(val2), val2)
```



## 优先级排序

()、[]  > ++/--/正负号/指针相关操作/! > 乘除，取余 > 加减  > 位运算 > 逻辑运算符 > +=/*=/^=等




