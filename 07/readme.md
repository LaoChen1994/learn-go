# 数组、切片、集合

## 数组的声明及赋值

**注意**：数组中可以是任意的类型，但是所有的元素必须都是相同的类型

**声明方式**：`var array_name [quantity]Type`

**赋值的方式**：`array[index] = Value`

一个输出素数结果的例子

```go
package main

import "fmt"

func main() {
    var array [4]int

    arrayIndex := 0
    flag := false

    for i := 2; i <= 10; i++ {
        if i == 2 {
            array[arrayIndex] = i
            arrayIndex++
            continue
        }
        for j := 2; j < i; j++ {
            if i%j == 0 {
                flag = true
                break
            }
        }

        if !flag {
            array[arrayIndex] = i
            arrayIndex++
        }

        flag = false
    }

    fmt.Println(array)
}
```

**缺点**：使用数组进行数据的存储问题在于，我们必须事先准确的声明数组的数量，但是很多时候我们计算不准就容易出现越界的问题，因此，Go中有另一种数据结构切片

## 切片的声明及赋值

**注意**：切片可以理解为没有“上限”的数组

**声明方式**：`var slice_name []Type`

**赋值方式**：`append`方法来对切片的值进行增加

**例子**

```go
package main

import "fmt"

func main() {
    var slice []int
    flag := false

    for i := 2; i <= 10; i++ {
        if i == 2 {
            slice = append(slice, i)
            continue
        }
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

    fmt.Println(slice)
}
```

## 集合的声明与赋值

**注意**：
    1. 集合的本质是键值对构成的哈希表，通过`key-value`的方式来进行存储
    2. 使用`new`关键用来申请一片内存空间，返回指向该内存空间的指针，`make`是用来创建`slice`、`map`、`Channel`等内部数据结构的

**声明方式**：`var map_name = make(map[key_type]value_type)`

**赋值方式**：`map_name[key] = Value`

**例子**

```go
func mapExample() {
    var stuInfo = make(map[string]string)

    stuInfo["001"] = "zhangSan"
    stuInfo["002"] = "liSi"
    stuInfo["003"] = "wangWu"

    fmt.Println(stuInfo) // map[001:zhangSan 002:liSi 003:wangWu]
}
```

## 获取切片的长度

**方法**：使用`len`这个内置的方法，`len([]int slice)`

**例子**：

```go
package main

import "fmt"

func main() {
	var slice []int
	flag := false

	for i := 2; i <= 10; i++ {
		if i == 2 {
			slice = append(slice, i)
			continue
		}
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
	fmt.Println(slice)
	fmt.Println("len ->", len(slice))
}
```



## 循环遍历

### 数组和集合的循环遍历

```go
func main () {
    var slice []int

    append(slice, 1)
    append(slice, 2)
    append(slice, 3)

    for j := 0; j < len(slice); j++ {
		slice[j] = slice[j] + 1
	}

	fmt.Println(slice)
}
```



### 集合的循环遍历

使用`range`和`for`进行遍历

```go
func mapExample() {
	var stuInfo = make(map[string]string)

	stuInfo["001"] = "zhangSan"
	stuInfo["002"] = "liSi"
	stuInfo["003"] = "wangWu"

	fmt.Println(stuInfo)

	for key, value := range stuInfo {
		fmt.Println("学号：", key, "姓名：", value)
	}
}
```


