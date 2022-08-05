# 并发（五）——定时器和Select

## 定时器

### Timer

**作用**：用于代码的延迟执行，类似`setTimeout`

**用法**：
    1. 借助`time`中的`Timer`类型
    2. `Timer`是一个结构体

```go
    type Timer struct {
        C <-chan Time
        r runTimer
    }
```

**解析**：看到上述的这个结构体应该能明白，在倒计时结束时，可以通过`<-Timer.C`来触发延迟任务的执行

**例子**:

这里有两个需要注意的点：
    1. `Timer`需要使用`time.NewTimer`创建出来
    2. 在结束时，需要延迟调用`delay.Stop`来防止内存泄漏

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	delay := time.NewTimer(2 * time.Second)
	fmt.Println("倒计时开始")
	defer delay.Stop()

	delaySeconds, ok := <-delay.C

	if ok {
		fmt.Println("倒计时结束", delaySeconds, "s")
	}
}
```

### Ticker

**作用**：用于代码的反复执行，类似`setInterval`
**例子**：使用方法如下

```go
ticker := time.NewTicker(2 * time.Second)
fmt.Println("每过2s常态化触发开始")
defer ticker.Stop()

end := 0

for val := range ticker.C {
    fmt.Println("有过了两秒啦~", val)

    end++

    if end > 10 {
        return
    }
}
```

## Select结构

**作用**：当一个接收者接收多个发送方传过来的数据，然后进行处理的场景

这里需要注意，如果需要判断两个任务同时结束，通过`sync.group`的方式需要判断标志位

```go

func syncTask() {
	wg.Add(3)
	dowloadProcrss := make(chan process)
	newDownload := make(chan int)
	go download(dowloadProcrss, newDownload)
	go downloadServer(dowloadProcrss)
	go addNewTask(newDownload)

	wg.Wait()
}

func download(ch1 chan process, ch2 chan int) {
	defer wg.Done()

	ch1Over := false
	ch2Over := false

	for {
		if ch1Over && ch2Over {
			return
		}

		select {
		case val, ok := <-ch1:
			if !ok {
				ch1Over = true
			}
			fmt.Println("1 -> ", val, ok)

		case val, ok := <-ch2:
			if !ok {
				ch2Over = true
			}
			fmt.Println("2 ->", val, ok)
		}
	}
}

func downloadServer(ch chan process) {
	defer wg.Done()
	defer close(ch)
	ch <- process{
		current: 0,
		total:   10,
	}
	for i := 0; i < 10; i++ {
		ch <- process{
			current: i + 1,
			total:   10,
		}
		time.Sleep(time.Second)

	}

}

func addNewTask(ch chan int) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	ch <- 1
	time.Sleep(2 * time.Second)
	ch <- 2
	close(ch)
}
```