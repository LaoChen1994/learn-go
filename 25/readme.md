# 并发（四）缓冲通道进行并发通信

## 缓冲通道的特点

1. 可以不必接消息和收消息一一匹配
2. 可以延迟接收消息，从存储的容器中去拿数据

## 示例代码

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

const DaysOfWeek = 7

func main() {
	wg.Add(2)
	intChan := make(chan int32, DaysOfWeek)
	go layEggs(intChan)
	go eatEggs(intChan)
	wg.Wait()
}

func layEggs(ch chan int32) {
	defer wg.Done()
	defer close(ch)

	var i int32 = 0

	for {
		if i < DaysOfWeek {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("放入1个鸡蛋")
			ch <- 1
			i++
		} else {
			return
		}
	}
}

func eatEggs(ch chan int32) {
	defer wg.Done()

	var count int32 = 0

	for {
		time.Sleep(1500 * time.Millisecond)
		val, ok := <-ch
		if ok && val > 0 {
			count += val
			fmt.Println("收到鸡蛋", val, "个", "，总计", count, "个")
		} else {
			return
		}
	}
}
```

## 构建安全的通道

如何构建安全的通道：多个协程之间的通信，由于协程之间都是异步操作，所以需要注意避免通道内存的泄露和死锁

**解决办法**：

1. 通道的关闭：只让某个通道的唯一发送者关闭该通道，因为 **关闭的通道，其他发送者无法再发送消息了**
2. 发送者最好自己单独启一条通道。
3. 不允许关闭一个已经关闭了的通道。
4. 接收者可以通过第二个参数来观察是否通道已经关闭`val, ok := <-ch`,`true`为通道开启，`false`为关闭

5. 使用`for range`语法可以简化获取通道参数的代码

```go
func collect(ch chan int32) {
	defer wg.Done()
	for val := range ch {
		fmt.Println(val)
	}
}
```

6. 如果一个通道中有一个生产者和两个消费者，两个香妃这的数据处于竞争关系, 这里的`collect`和`eatEggs`会共同消费`layEggs`产生的鸡蛋数据

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

const DaysOfWeek = 7

func main() {
	wg.Add(3)
	intChan := make(chan int32, DaysOfWeek)
	go layEggs(intChan)
	go eatEggs(intChan)
	go collect(intChan)
	wg.Wait()
}

func layEggs(ch chan int32) {
	defer wg.Done()
	defer close(ch)

	var i int32 = 0

	for {
		if i < DaysOfWeek {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("放入1个鸡蛋")
			ch <- 1
			i++
		} else {
			return
		}
	}
}

func eatEggs(ch chan int32) {
	defer wg.Done()

	var count int32 = 0

	for {
		time.Sleep(1500 * time.Millisecond)
		val, ok := <-ch
		if ok && val > 0 {
			count += val
			fmt.Println("收到鸡蛋", val, "个", "，总计", count, "个")
		} else {
			return
		}
	}
}

func collect(ch chan int32) {
	defer wg.Done()
	for val := range ch {
		fmt.Println(val)
	}
}
```

## 单向通道的创建

1. 接收通道
```go
var readOnlyIntChan <-chan int = intChan
```

2. 只发送通道
```go
var sendOnlyIntChan chan<- int = intChan

```