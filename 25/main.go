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
