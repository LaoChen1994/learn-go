# 开篇介绍

## Go语言优势

1. 相较于Java容易上手

2. 具有更好的并发编程优势和自带的标准库

3. Go SDK中自带的format完成代码格式化



## GO语言知识图谱

![小册知识点结构.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/4a076b7d583e435c83acd7515d741928~tplv-k3u1fbpfcp-zoom-in-crop-mark:3024:0:0:0.awebp?)



## GO语言的安装

1. 官网下载对应的包[Downloads - The Go Programming Language](https://golang.google.cn/dl/)

2. 配置环境变量
   
   1. GOPATH：执行go命令时查询的目录
      
      windows：`%USERPROFILE%\go`
      
      linux/mac: `$home/go`
   
   2. GOROOT：表示GO安装的目录，多个不同的GO的SDK之间的版本切换通过GOROOT来做切换
   
   3. GOBIN：做GOROOT做别的版本切换的时候，用`GOROOT/bin`来标记对应编译后的二进制文件，来进行版本的切换
      
      

## Hello GO

1. 编写第一个GO的程序

```go
package main
import "fmt"

func main() {
	fmt.Println("Hello World");
}
```

2. 第一次编译

```bash
go build hello.go
```

3. 编译完成后会生成一个`hello.exe`文件，直接在命令行中运行`./hello`即可执行对应的文件


