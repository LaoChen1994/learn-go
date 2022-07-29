# 文件操作

## 目标

目标：实现一个从命令行获取参数，最终保存到文件的过程

## 获取命令行参数

1. 使用`os.Args`这个方法即可
2. 例子代码

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)
}
```

## 获取文件信息

### 用到的包

创建和获取文件信息需要用到的包和函数：
+ 包：os
+ 相关方法：`Stat`,`IsNotExist`,`Create`

### 创建流程
+ 先获取文件
+ 文件不存在则创建文件
+ 如果文件存在且是个文件夹就抛出错误

```go
func createFile(filePath string) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(filePath)

			if err != nil {
				fmt.Println("创建文件失败, ", err)
				panic(err)
			}

			file.Close()
			return
		} else {
			panic(err)
		}
	}

	if fileInfo.IsDir() {
		panic("已存在文件为文件夹，请确认路径")
	}
}
```

### 读取文件

**具体步骤**：
1. 打开文件，这里有三个参数
   1. 文件地址
   2. 打开的方式，只读/只写/读写等，参见os的枚举
   3. 权限：可读可写可执行
2. 直接用`file.Write`即可

```go
func main() {
	args := os.Args[1:]
	var filePath = "E:/Learn/learn-go/21/1.txt"

	file, err := os.OpenFile(filePath, os.O_APPEND, os.ModePerm)

	if err != nil {
		panic(err)
	}

	str := strings.Join(args, " ") + "\n"
	file.Write([]byte(str))
	file.Close()
}

```


### 读取文件

使用`ioutil`来进行文件的读取

例子：

```go
content, error := ioutil.ReadFile(filePath)

if error != nil {
	panic(error)
}

val := string(content)

fmt.Println(val)

```