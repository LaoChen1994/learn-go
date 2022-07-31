# 并发（一）——GO语言并发初探

## 基本概念

### 并发与并行

并发：在一个CPU上通过时间分片执行不同的任务，同时运行多个任务
并行：多个CPU同时执行多个任务

### 线程和协程

进程：一个应用占用一个进程，在一个进程中可以有多个线程来执行操作，进程一旦终止，进程中的线程和协程也将终止

线程：线程是操作系统的资源，创建、切换、停止由操作系统控制

协程：创建、切换停止由用户操作，更轻量化，一个线程上可以跑多个协程，**在GO语言中，并发通过协程来实现，大量IO密集型的操作适合协程进行处理**

## 并发任务的启动

启动方式：在对应的执行语句之前增加`go`关键字即可

**注意点**：因为我们不知道什么时候协程会执行结束，这里简单粗暴的通过一个延迟来让其能够一直运行到我们期望的`testGoFunc`结束
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go testGoFunc(1)
	go testGoFunc(2)
	time.Sleep(time.Second * 10)

	fmt.Println("结束")
}

func testGoFunc(delay int32) {
	for i := 0; i < 10; i++ {
		fmt.Printf("当前是%d延迟，第%d次运行\n", delay, i)
		time.Sleep(time.Second * time.Duration(delay))
	}
}
```

### sync的使用

上述代码中通过一个延迟时间来保证协程能够执行结束，在实际工作中，无法预想协程内的程序多久能执行完毕，GO中可以使用`sync`这个包来保证

**使用方法**：
1. 声明一个`sync.WaitGroup`的变量
2. 在主线程中调用`goRoutineWait.Add`声明需要开启几个`GoRoutine`
3. 在协程中通过`goRoutineWait.Done()`来标记该协程完成
4. 在主线程中通过`goRoutinueWait.Wait()`表示在这里一直等到协程执行结束

**例子**：

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var goRoutineWait sync.WaitGroup

func main() {
	fmt.Println("开始")
	goRoutineWait.Add(2)

	defer fmt.Println("结束")
	defer goRoutineWait.Wait()

	go testGoFunc(1)
	go testGoFunc(2)

}

func testGoFunc(delay int32) {
	defer goRoutineWait.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("当前是%d延迟，第%d次运行\n", delay, i)
		time.Sleep(time.Second * time.Duration(delay))
	}
}
```

如果上述的协程只运行一次，也可以通过匿名函数IIFE来执行，这样会使代码更加精简