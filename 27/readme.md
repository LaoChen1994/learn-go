# 并发（六）——锁和原子操作

为什么需要**锁和原子操作**，因为并发是不安全的，观察下面的代码，理论上两个并发任务需要各加1000，结果为2000，
但是其实结果会远远小于2000，这是因为共享的内存数据`res`存在数据竞争，被两个协程同时抢夺

```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var res int = 0

func main() {
	wg.Add(2)
	go add()
	go add()

	wg.Wait()
	fmt.Println(res)
}

func add() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		res++
	}
}
```

## 锁

### 互斥锁

**什么是互斥锁**：被加锁的代码一旦被运行，必须等到解锁后才能继续运行，**这里加锁的是运行的代码**

**如何申明互斥锁**：`var mutexLocker sync.Mutex`

**例子**：

```go
var mutexLocker sync.Mutex

func add() {
	defer wg.Done()
    // 执行完后解锁
	defer mutexLocker.Unlock()
    // 加锁
	mutexLocker.Lock()
	for i := 0; i < 1000; i++ {
		res++
	}
}

```

**缺点**：使用这种互斥锁的话，就相当于将并行的并发任务因为锁的存在直接变成了串行任务，就失去了并发的优势

### 读写互斥锁

**什么是读写互斥锁**：理论上只要共享的数据不发生改变，几乎不会用到锁，对读数据的操作加锁，反而影响性能，**读写互斥锁的目标是在写操作上串行，保证数据安全**

**声明读写互斥锁**：`var lock sync.RWMutex`

**用法**：使用读锁`sync.Rlock()`和`sync.RUnlock()`,写锁依然是`sync.lock()`和`sync.unlock()`

**例子**：

这里需要关注的点是：
    1. 写锁执行的时候所有并行任务会转换为串行任务
    2. **如果当前是写操作，读锁和写锁都需要等待**
    3. 写操作需要等待已运行的读操作锁结束后才开始运行
    4. 加锁的时候避免A等B, B等C，C等A导致的死锁 

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var money float32 = 100.00
var wg sync.WaitGroup
var lock sync.RWMutex

func main() {
	fmt.Println("start =>", time.Now().Format("15:04:05"))
	wg.Add(6)
	go readMoney(5)
	go readMoney(2)
	go saveMoney(100.00)
	go readMoney(4)
	go saveMoney(200.00)
	go readMoney(1)
	wg.Wait()

	fmt.Println("end =>", time.Now().Format("15:04:05"))

}

func saveMoney(m float32) {
	defer wg.Done()
	defer lock.Unlock()
	lock.Lock()

	time.Sleep(1 * time.Second)
	money += m
	fmt.Println("money 存入", m, "总计:", money)
}

func readMoney(duration int32) {
	defer wg.Done()
	defer lock.RUnlock()
	lock.RLock()

	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Println(time.Now().Format("15:04:05"), "=> duration =>", duration, "money => ", money)
}
```

## 原子操作

**什么是原子操作**：就是指那些在运行过程中不能被打断的操作

**如何实现**：可以使用`sync/atom`这个包来实现原子操作，原子操作也会使两次任务串行化，是直接通过CPU指令来进行操作的方式，性能较锁更快

**原子操作的特性**：
+ 原子操作都是非入侵式的；
+ 原子操作共有五种：增减、比较并交换、载入、存储、交换；
+ 原子操作支持的类型类型包括 int32、int64、uint32、uint64、uintptr、unsafe.Pointer


**例子**

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var res int32 = 0

func main() {
	wg.Add(2)
	go add1()
	go add2()

	wg.Wait()
	fmt.Println(res)
}

func add1() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(&res, 1)
	}
}

func add2() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(&res, 2)
	}
}
```