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
