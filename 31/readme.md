# 模块化

## 模块化步骤

### 使用GO MOD来做依赖管理

1. 开启GO111MODULE
   ```go
    go env -w GO111MODULE=on
   ```
2. 使用`go mod`初始化项目
   ```go
    go mod init pd
   ```

### 编写自定义模块

因为之前`int`的module是pd，所以所有模块的引入必须是`pd`开头，编写一个简单的`add`模块

```go
// 文件位置为 add/add.go
package add

func Add(num1 int32, num2 int32) int32 {
	return num1 + num2
}
```

### 引入模块

```go
package main

import (
	"fmt"
	"pd/add"
)

func main() {
	rlt := add.Add(1, 2)

	fmt.Println(rlt)
}
```