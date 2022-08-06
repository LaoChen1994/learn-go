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
