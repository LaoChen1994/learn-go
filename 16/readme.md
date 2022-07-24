# 接口（三）——灵活安全地使用接口

## 接口的嵌套组合

**方法**：将多个`interface`作为需要嵌套接口的结构体即可

**定义例子**：

幻想一个文件下载的例子，选择文件地址和保存，我们将其拆分成两个接口，使用`FileDownload`来进行聚合

```go
package main

import "fmt"

type FileDest interface {
	chooseDest(url string, name string) string
}

type Save interface {
	saveFile()
}

type FileDownload interface {
	FileDest
	Save
}

type File struct {
	path string
	name string
}

func (f *File) chooseDest(url string, name string) string {
	f.path = url
	f.name = name
	return url
}

func (f *File) saveFile() {
	fmt.Println(f.name, "存储成功, 地址", f.path)
}
```

**聚合接口的使用**：

在使用聚合的接口他内部的方法会进行平铺放到第一层，所以通过结构体直接调用即可

```go

func main() {
	var downloader FileDownload

	downloader = new(File)
	downloader.chooseDest("/user/download/1.txt", "1.txt")
	downloader.saveFile()

}

```

## 从空接口中取值

**方法**：使用`x.(type)`这种方案对空接口的类型进行判断可能更好

**例子**：

```go
func dataOutput(data interface{}) {
	dataVal, stringOk := data.(string)

	if stringOk {
		fmt.Println(dataVal)
	}
}
```

## 空接口的值比较

**注意**：
    1. 在GO语言中不可比较的类型`Map`和`Slice`，上述两种类型只能和`nil`进行比较
    2. 可以比较的类型：基础类型和`数组`

**例子**：

```go
var a = [5]int{1, 2, 3, 4, 5}
var b = [5]int{1, 2, 3, 4, 5}

// true
fmt.Println(a == b)
```

```go
c = []int{1, 2, 3, 4}
d = []int{1, 2, 3, 4}

// invalid operation: cannot compare c == d (slice can only be compared to nil)
fmt.Println(c == d)

```

## 空接口与nil

**作用**：`nil`可以赋值给指针类型和接口类型，用于给一个默认值，来判断是否这个指针类型为空

**例子**

```go
func getInitLoader() *FileDownload {
	var downloader *FileDownload = nil

	return downloader
}

func main() {
    // nil
	fmt.Println(getInitLoader())
    // true
	fmt.Println(getInitLoader() == nil)
}
```