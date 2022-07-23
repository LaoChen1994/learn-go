# 接口（二）——泛型和断言

## 空接口

**什么是泛型**：数据类型中的万能类型，可以存放任何类型的数据，GO语言中通过 **空接口** 来实现

在`go` 1.17版本下，可以使用空接口来实现泛型，其类似于`typescript`中的`any`类型

**例子**

```go
package main

import "fmt"

func paramBack(value interface{}) interface{} {
	return value
}

func main() {
	fmt.Println(paramBack(123))
	fmt.Println(paramBack(123.33))
	fmt.Println(paramBack("123aaa"))
}
```

## 泛型的定义和使用

**泛型定义**：在`go` 1.18版本以后可以通过`type`来申明指定的类型，用`T`来代替指定的类型

**例子**：
```go
// 定义一个具有string int float 的切片
type SliceType[T string | int32 | float64] []T
type SliceMap[T string | int32] map[string]SliceType[T]
type StructType[T int | string] struct {
	name   string
	salary T
}
```

**注意**：
	1. 基础类型能只有类型参数，比如`type AnyType[T string | int32 | float64] []` 是错误语法
	2. 对于引用类型需要通过`interface`进行包装，避免编译器误解`type NewType2[T interface{*int|*float64}] []T `
	3. 通过泛型的嵌套可以实现各种嵌套复杂的数据结构
	4. 匿名结构体不支持泛型

**泛型的使用**

泛型的调用方法，是通过为对应的`type`添加方法的方式，并通过调用方法来实现对应的功能

```go
type NumberType[T int32 | float32] []T
func (s NumberType[T]) add() T {
	var rlt T
	for _, v := range s {
		rlt += v
	}

	return rlt
}

func main()  {
	var sliceInt NumberType[int32]
	sliceInt = append(sliceInt, 1)
	sliceInt = append(sliceInt, 2)
	sliceInt = append(sliceInt, 3)

	var sliceFloat NumberType[float32]

	sliceFloat = append(sliceFloat, 1.1)
	sliceFloat = append(sliceFloat, 1.2)
	sliceFloat = append(sliceFloat, 1.3)

	fmt.Println(sliceInt.add())
	fmt.Println(sliceFloat.add())
}
```

**泛型函数用法**：上述通过`receiver`的方式来进行调用是不是感觉有点繁琐，除此以外，我们还可以直接使用泛型函数来进行声明

```go
func Sum[T int32 | float32](list NumberType[T]) T {
	var rlt T
	for _, v := range list {
		rlt += v
	}

	return rlt
}

func main()  {
	var sliceInt NumberType[int32]
	sliceInt = append(sliceInt, 1)
	sliceInt = append(sliceInt, 2)
	sliceInt = append(sliceInt, 3)

	res := Sum(sliceInt)
	fmt.Println(res)
}

```

## 类型断言

这里注意一点就是，空接口是可以是代表任何类型，如何使用类型判断调用不同的方法呢？可以使用`value, ok = x.(T)`这个方式来对类型进行判断

```go
type Cube[T int32 | float32] struct {
	length T
}

type Cuboid[T int32 | float32] struct {
	width  T
	length T
	height T
}

func (c Cube[T]) calcVolum() T {
	return c.length * c.length * c.length
}

func (c Cuboid[T]) calcVolum() T {
	return c.length * c.width * c.height
}

func calcVolum[T int32 | float32](thing interface{}) T {
	cube, ok := thing.(Cube[T])

	if ok {
		return cube.calcVolum()
	} else {
		return 0
	}
}

func main()  {
	cube := Cube[int32]{
		length: 3,
	}

	volumn := calcVolum[int32](cube)

	fmt.Println(volumn)
}
```

