# 构建“轮子”——函数

## 函数定义和调用

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

## 值传递和引用

**值传递**：**直接将值传递到函数中，会直接在函数内部生成一个副本**，所以对应内存中是不同的地址，内部参数的变化不会影响外部。

**引用传递**：传递的是同一个地址，所以其本身内部值的改变会影响外部的值

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


