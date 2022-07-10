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
