# GO程序源码结构、注释、命令行工具

## Hello Go

```go
// 表示这个源码文件属于main包
// 每个go应用程序有且仅有一个main包
package main

// 导入fmt的包，之后可以使用fmt这个包的各种能力
import "fmt"

func main() {
	/*
		调用fmt包中的Println来输出文案
	*/
	fmt.Println("Hello world")
}


```



这里使用`//`来注释单行内容，使用`/**/`来注释多行文字



## Go SDK命令行工具

### go build

作用：用于编译GO的源码

优点：支持并发编译

支持参数：

+ -v：编译时显示包名

+ -p x：指定编译时的核心数为`x`

+ -a：强制进行重新构建

+ -n：仅输出编译时执行的所有命令

+ -x：执行编译并输出所有命令

+ -race：开启静态检测



### go clean

作用：清除当前目录下生成的与包名或者go源码同名的可执行文件

参数：

+ -i：清除关联安装的包和可运行文件

+ -n：仅输出清理时执行的所有命令

+ -r：递归清除在import中引入的包

+ -x：执行清理并输出清理时执行的所有命令

+ -cache：清除缓存



### go run

作用：直接执行go的源码且不生成可执行文件



### gofmt

作用：用于格式化go的代码

具体参数参照`gofmt --help`



### go install

作用：和`go build`类似，将源码编译为可执行文件

区别：

+ 编译源码后，将可执行文件或库安装到约定的目录下

+ 生成的可执行文件用包名来命名

+ 默认情况下会安装到`GOPATH/bin`目录下



### go get

作用：用于安装第三方包，下载源码后并执行`go install`。

具体参数：

+ -d：仅下载包，不安装

+ -f：不验证每个包是否被导入

+ -fix：下载后执行fix操作

+ -u：更新源码到最新版本

+ -u=patch：小版本更新到最新

+ -v：执行获取并实时显示日志



例子：

```bash
go get github.com/ethereum/go-ethereum
# 获取指定版本
go get github.com/ethereum/go-ethereum@v1.10.1
```
