# 流程控制
## 循环结构

1. 循环结构格式
   
   `for init; condition;post {}`

   `for _, v := range s {}`

2. 一个杨辉三角形的例子
   
   ```go
   package main
   
   import "fmt"
   
   func main() {
   	n := 7
   
   	for i := 1; i <= n; i++ {
   		for j := 0; j < n-i; j++ {
   			fmt.Print(" ")
   		}
   
   		for k := 0; k < 2*i-1; k++ {
   			fmt.Print("*")
   		}
   
   		fmt.Println("")
   	}
   
   	for i := n - 1; i > 0; i-- {
   		for j := 0; j < n-i; j++ {
   			fmt.Print(" ")
   		}
   
   		for k := 0; k < 2*i-1; k++ {
   			fmt.Print("*")
   		}
   
   		fmt.Println("")
   	}
   }
   
   ```

## 条件控制if

1. go中的流程控制语句模板
   
   ```go
   if condition1 {
       
   }else if condition2 {
       
   } else {
       
   }
   ```

2. 一个查找10以内素数的例子
   
   ```go
   func findNum() {
   	flag := false
   
   	for num := 2; num < 10; num++ {
   		if num == 2 {
   			fmt.Println(num)
   			continue
   		} else {
   			for i := 2; i < num; i++ {
   				if num%i == 0 {
   					flag = true
   					break
   				}
   			}
   		}
   
   		if !flag {
   			fmt.Println(num)
   		}
   
   		flag = false
   	}
   }
   ```

3. 使用break和continue可以中断和跳过`for`循环

## 条件控制switch

和其他语言中一样，go中也有switch-case的用法

1. go中switch用法

```go
	switch num {
		case 0:
		case 1:
			f()
		case 2:
			f2()
	}

```

2. 例子

在go中当匹配到一个case以后直接退出整个`switch`块, 不会再进入到其他的分支逻辑中

如果我们希望继续能走到其他的运行逻辑，可以使用，`fallthrough`关键字来让其继续走到后续的case中

下面的例子中，新用户进来就会附赠一句欢迎光临

```go
func switchFunc(num int32) {
	switch num {
	case 2:
		fmt.Println("这是老用户了")
	case 1:
		fmt.Println("这是第一次来的用户")
		fallthrough
	default:
		fmt.Println("欢迎光临")
	}
}
```