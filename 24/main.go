package main

import (
	"fmt"
	"sync"
)

var intChan = make(chan int)
var syncWait sync.WaitGroup

func main() {
	syncWait.Add(2)

	go layEggs()
	go eatEggs()
	syncWait.Wait()
}

func layEggs() {
	left := 5
	defer syncWait.Done()
	intChan <- 1
	left--
	for {
		_, ok := <-intChan

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
		fmt.Println("eat ->", eggsCounts)
		if !ok {
			fmt.Println("我也吃饱了")
			return
		}
		intChan <- 0
	}
}
