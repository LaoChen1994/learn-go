# 并发（二）——协程间通信

## CSP并发模型

并发通信的方式：
    1. 通过多线程任务`内存共享`，通过锁的方式来保证数据安全(**内存共享**)
    2. 通过CSP模型，通过`通信方式共享`内存数据(**通信共享**)

GO并发的三个陷阱：
    1. 协程任务泄露（运行失控导致一直留在内存中，引起内存泄漏）
    2. 未完成的任务（任务没有执行完成）
    3. 数据竞争


## 通信Channel类型

GO语言中的通道：
    1. 同步通道（A协程依赖B协程的数据，B等待A的数据再执行，类似js的`await`）
    2. 缓冲通道（生产消费模式，直接把数据放到缓冲区，消费者消费）

## 同步通道

**通道的创建**：使用`make`关键字，`make(chan type)`来创建一个类型为`type`的Channel

**通道的赋值和接收**：箭头操作符

```go
// 放入数据
intChan <- 1

// 取出数据
val := <-intChan
```

**例子代码**

```go
package main

import (
	"fmt"
	"reflect"
	"sync"
)

var intChan = make(chan int)
var syncWait sync.WaitGroup

func main() {
	fmt.Println(intChan)
	fmt.Println(reflect.TypeOf(intChan))

	syncWait.Add(2)

	go layEggs()
	go eatEggs()
	syncWait.Wait()
}

func cc(str string) {
	fmt.Println(str)
}

func layEggs() {
	defer syncWait.Done()
	defer close(intChan)
	intChan <- 1
	fmt.Println("lay egg")
}

func eatEggs() {
	defer syncWait.Done()
	eggsCounts := <-intChan
	fmt.Println("Eat ", eggsCounts, "eggs")
}

```

## 多次通信的场景

如果来两个线程来回进行通信，同步通道当对方的通道没有发送消息，另外一个通道会一直等到陷入死锁，需要注意

```go
package main

import (
	"fmt"
	"reflect"
	"sync"
)

var intChan = make(chan int)
var syncWait sync.WaitGroup

func main() {
	fmt.Println(intChan)
	fmt.Println(reflect.TypeOf(intChan))

	syncWait.Add(2)

	go layEggs()
	go eatEggs()
	syncWait.Wait()
}

func cc(str string) {
	fmt.Println(str)
}

func layEggs() {
	left := 4
	defer syncWait.Done()
	intChan <- 1
	for {
		_, ok := <-intChan
		fmt.Println("lay ->", ok)

		if left <= 0 || !ok {
			fmt.Println("下不出蛋了")
			close(intChan)
			return
		}
		intChan <- 1
		fmt.Println("lay egg")
		left--
	}
}

func eatEggs() {
	defer syncWait.Done()
	for {
		eggsCounts, ok := <-intChan
		fmt.Println("eat ->", eggsCounts, ok)
		if !ok {
			fmt.Println("我也吃饱了")
			return
		}
		fmt.Println("Eat ", eggsCounts, "eggs")
		intChan <- 0
	}
}
```