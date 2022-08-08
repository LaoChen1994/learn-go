# 反射（二）结构体、函数、实例

## 使用反射访问结构体

### 遍历结构体

方法：使用`ValueOf.Field`来遍历对应字段的key

**注意**：如果要获取对应的类型相关的参数可以遍历`Type.Field`

例子：

```go
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int
	Gender int
}

func main() {
	personExample := Person{
		Name:   "小王",
		Age:    18,
		Gender: 1,
	}

	personType := reflect.TypeOf(personExample)

	fmt.Println(personType.Name())
	fmt.Println(personType.Kind())

	fmt.Println(personType.NumField())

	values := reflect.ValueOf(personExample)
	for i := 0; i < values.NumField(); i++ {
		field := values.Field(i)
		name := personType.Field(i).Name
		t := personType.Field(i).Type

		fmt.Println(name, t, field)
	}

}
```

### 获取特定字段的值和设置对应的值

方法：使用`FieldByName`来获取对应的值，兵通过`SetInt`等方式来进行设置值, 需要注意的是，引用类型才可赋值,但是之前代码中的`NumField`需要针对值引用才可用，所以用`Elem()`将引用指针变成值引用

例子：
```go
    personExample := &Person{
		Name:   "小王",
		Age:    18,
		Gender: 1,
	}
    values := reflect.ValueOf(personExample).Elem()
	fmt.Println("age =>", values.FieldByName("Age"))
	values.FieldByName("Name").SetString("小红")
	values.FieldByName("Age").SetInt(20)
	fmt.Println(values.FieldByName("Name").IsValid(), values.FieldByName("Name").CanSet())
	fmt.Println("name =>", values.FieldByName("Name"))
	fmt.Println("age =>", values.FieldByName("Age"))
```

**注意点**：这里需要注意的是，可以通过反射修改的值有这么几个特点：1.可寻址 2. 可导出（结构体的参数名为大写的表示可导出）

## 使用反射调用函数

方法：使用`reflect.ValueOf`获取函数值，之后使用使用`relect.valueOf`来构造参数，通过`Call`来对函数进行调用

例子：

```go
func addCalc(num1 int, num2 int) int {
	return num1 + num2
}

func main()  {
	accFuncType := reflect.TypeOf(addCalc)
    // func
	fmt.Println(accFuncType.Kind())

    // 构造参数
	accFuncParam := []reflect.Value{reflect.ValueOf(20), reflect.ValueOf(200)}

	accValue := reflect.ValueOf(addCalc)

    // 用call方法来调用函数
	res := accValue.Call(accFuncParam)

	fmt.Println(res[0])    
}

```

## 使用反射创建实例

方法：使用`reflect.New`来创建一个和指定变量同类型的变量

例子：

```go
var num myint = 100

typeOfNum := reflect.TypeOf(num)
another := reflect.New(typeOfNum)

another.Elem().SetInt(200)

fmt.Println(another.Elem().Int())
fmt.Println(num)
fmt.Println(another.Type().Kind(), typeOfNum.Kind())
```