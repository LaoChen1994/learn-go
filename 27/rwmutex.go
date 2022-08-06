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
	go readMoney(2)
	go readMoney(5)
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
	defer fmt.Println("存入结束 =>", time.Now().Format("15:04:05"))

	lock.Lock()

	time.Sleep(1 * time.Second)
	fmt.Println("存入开始 =>", time.Now().Format("15:04:05"))
	money += m
	fmt.Println("money 存入", m, "总计:", money)
}

func readMoney(duration int32) {
	defer wg.Done()
	defer lock.RUnlock()
	defer fmt.Println(duration, "读取结束 =>", time.Now().Format("15:04:05"))
	lock.RLock()

	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Println(duration, "读取开始 =>", time.Now().Format("15:04:05"))
	fmt.Println(time.Now().Format("15:04:05"), "=> duration =>", duration, "money => ", money)
}
