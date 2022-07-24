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

## 条件控制与流程控制

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
