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

**注意**：集合的本质是键值对构成的哈希表，通过`key-value`的方式来进行存储

**声明方式**：`var map_name = make(map[key_type]value_type)`

**赋值方式**：`map_name[key] = Value`

**例子**

```go

```


