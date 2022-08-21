# 构建“轮子”——函数

## 函数定义和调用

几个注意点：
1. `main`和`init`函数不可以有参数和返回值
2. 在GO中函数的重载是不被允许的

**函数的定义方式**：

```go
func function_name([params_list])[return_value_list]{}
```

**一个字符串拼接的例子**

```go
package main

import (
    "fmt"
    "strings"
)

func strFunc(val []string) string {
    return strings.Join(val, ",")
}

func main() {
    var list []string
    list = append(list, "123")
    list = append(list, "456")
    res := strFunc(list)

    fmt.Println(res)
}
```

**先置函数返回声明的方式**

在函数返回值的时候，可以在声明返回值的时候就先声明需要返回值变量的名称，这样在后续就不需要返回对应具体的变量，**注意的是`return`必须要**

```go
func main() {
	x2, x3 := getX2X3(4)

	fmt.Println(x2, x3)
}

func getX2X3(num int) (x2 int, x3 int) {
	x2 = num * 2
	x3 = num * 3
	return
}

func mult_returnval(a int32, b int32) (sum int32, diff int32) {
	sum = a + b
	diff = a - b
	return
}
```

**函数的声明**

在GO中编写函数不一定一定要带上函数体，可以只做函数的声明

```go

func add (a, b int) int

```

## 值传递和引用

**值传递**：**直接将值传递到函数中，会直接在函数内部生成一个副本**，所以对应内存中是不同的地址，内部参数的变化不会影响外部。

**引用传递**：传递的是同一个地址，所以其本身内部值的改变会影响外部的值, **在类似slice, map, interface, channel**这些引用类型传递时，默认都是使用引用传递，不管是否显示指定了指针传值`&`

**值引用的例子**

```go
func main() {
    var rlt []int
    findPrimeNumber(rlt, 100)
    // 这个时候rlt仍为[],这个时候是值引用
    fmt.Println(rlt)
}


func findPrimeNumber(slice []int, max int) []int {
    for i := 1; i < max; i++ {
        if i == 2 {
            slice = append(slice, i)
            continue
        }

        flag := false

        for j := 2; j < i; j++ {
            if i%j == 0 {
                flag = true
                break
            }
        }

        if !flag {
            slice = append(slice, i)
        }

        flag = false
    }
}
```

**引用传递的例子**

参数中通过`*`来做标记来代表这个参数是**引用的地址**，这里需要和`&`用来取地址和`*&`用来取地址内的值的进行区分开

```go
func main() {
    var rlt []int

    findPrimeNumber(&rlt, 100)
    fmt.Println(rlt)
}

func findPrimeNumber(slice *[]int, max int) {
    for i := 1; i < max; i++ {
        if i == 2 {
            *slice = append(*slice, i)
            continue
        }

        flag := false

        for j := 2; j < i; j++ {
            if i%j == 0 {
                flag = true
                break
            }
        }

        if !flag {
            *slice = append(*slice, i)
        }

        flag = false
    }
}
```

## 缺省参数的用法

可以使用 `...type` 的方式来获取无限长度的`slice`类型的缺省参数，缺省参数可以作为`slice`进行二次传递

```go
func min(a ...int) (val int) {
	val = 0

	if len(a) == 0 {
		return
	}

	val = a[0]

	for _, v := range a {
		if v < val {
			val = v
		}
	}

	return
}

// 用法

fmt.Println(min(8, 2, 9, 4, 5, 1))
```

## 函数的延迟调用

**用法**：使用`defer`内置关键词，当代码执行到这里的时候会延迟调用，`defer`申明的越前面，则执行的顺序在越后面

**例子**

```go
func main() {
	defer fmt.Println("我是最后一个出来的")
	defer fmt.Println("我是倒数第二个出来的")
	defer fmt.Println("我是倒数第三个出来的")

	var list []string
	list = append(list, "123")
	list = append(list, "456")
	res := strFunc(list)

	var rlt []int

	findPrimeNumber(&rlt, 100)

	fmt.Println(res)
	fmt.Println(rlt)
}
```



## 递归算法

**注意**：算法中自己调用自己和终止条件

**例子**

```go
func calcFactorial(n int) (result int) {
	if n > 0 {
		return n * calcFactorial(n-1)
	}

	return 0
}
```

```go
func calcFibonacci(n int) (result int) {
	if n == 1 || n == 2 {
		return 1
	}

	return calcFibonacci(n-1) + calcFibonacci(n-2)
}
```


