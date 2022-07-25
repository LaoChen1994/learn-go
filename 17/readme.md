# 包（一）——GO语言是如何跑起来的

## 包的声明

先看一个简单的程序:

```go
package main

import "fmt"

func init() {
	fmt.Println("我在 main 之前输出")
}

func main() {
	fmt.Println("Hello world")
}
```

从上述的程序中我们可以看到以下几点：

1. 上述程序中`package main`声明了当前的包为`main`这个是一个特殊的包名，**每个GO的程序必须拥有一个main**的包
2. `main`函数必须放在`main`包中，，如果想要编译出可执行文件，必须有`main`函数，不然编译会报错
3. 除了`main`以外，还有一个`init`的初始函数，将先于`main`执行
4. 使用`import`关键字来引入其他的包

## GO源码的启动流程

[GO源码启动流程](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/6c91ccc186824ea8b8ad06dc2b28faf8~tplv-k3u1fbpfcp-zoom-in-crop-mark:3024:0:0:0.awebp?)

大致过程为：
1. 从`main`包开始执行
2. 导入依赖包，如果依赖又导入了其他依赖会深度遍历向下递归导入，知道导入依赖完毕后才会执行本包的后续代码