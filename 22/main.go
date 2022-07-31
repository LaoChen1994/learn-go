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
