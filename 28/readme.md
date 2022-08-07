# 反射（一）基础

## 反射的基本使用

**为什么要用反射**：可以使程序在运行期间对其进行访问和修改，**在编译完成后无法获取变量名、变量类型以及结构体内部构造信息的**

Go中使用`reflect`这个包

### 获取变量的类型和值

使用`reflect.TypeOf`和`reflect.ValueOf`来实现

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	testNum := 10

    // int
	fmt.Println(reflect.TypeOf(testNum))
    // 10
	fmt.Println(reflect.ValueOf(testNum))
}
```

**注意点**：
1. 通过`ValueOf`获得的值是`Value`类型，如果需要获取指定类型的值，需要通过类型转换实现，但是因为是在运行时操作的，所以类型异常在编译时无法被识别，这个时候会引发宕机

```go
testNum := 10
value := reflect.ValueOf(testNum)
// 这种方式是直接对值的操作
val := int32(value.Int())
fmt.Println(val)
```

2. 判断值是否有效的方法是`isNil`和`isValid`

```go
fmt.Println(value.IsValid())
fmt.Println(value.IsNil())
```

3. 在操作反射时，必须要保证操作对象是可寻址的，因为函数间值的传递都是指引用，我们通过反射是直接改动 **地址内对应的值**，而不是传参值的副本

### 针对引用值的获取

通过注意点3中我们知道，反射需要修改时需要对引用指针进行操作，如果直接通过`Value.int()`来获取值，这个时候会报错，因为这个时候的`Value`是一个指针，需要通过`Elem`来获取对应的值

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var testNum int = 10
	testNumPtr := &testNum
	value := reflect.ValueOf(testNumPtr)

    // 通过Elem来从指针中取值
    fmt.Println(value.Elem().CanSet(), value.Elem().CanAddr())


	if value.IsNil() || !value.IsValid() {
		return
	}

	fmt.Println(value.Elem().Int())
}
```

## 通过引用类型来修改值

使用`Elem().SetInt()`方法即可

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var testNum int = 10
	testNumPtr := &testNum
	value := reflect.ValueOf(testNumPtr)

    // 10 10 10
	fmt.Println(value.Elem().Int(), testNum, *testNumPtr)

	value.Elem().SetInt(20)

    // 20 20 20
	fmt.Println(value.Elem().Int(), testNum, *testNumPtr)
}
```

## 反射的三大定律


1. 反射可以将接口类型变量转换成反射类型对象，使用`ValueOf`和`TypeOf`将变量先转成`interface{}`之后再转换成`Value`和`Type`的反射类型

2. 将反射类型对象转换成接口类型变量，使用`reflect.Value.interface()`将`Value`转换成任意类型，之后用`.Int`等方法转换成其他变量类型

3. 如果要修改“反射类型对象”，其值必须是“可写的”(CanSet)

```go
	var testNum int = 10
	num := reflect.ValueOf(testNum)

    // false
	fmt.Println(num.CanSet())
	if num.CanSet() {
		num.SetInt(30)
	}
```