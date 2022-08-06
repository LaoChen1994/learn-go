package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutexLocker sync.Mutex
var res int = 0

func main() {
	wg.Add(2)
	go add1()
	go add2()

	wg.Wait()
	fmt.Println(res)
}

func add1() {
	defer wg.Done()
	defer mutexLocker.Unlock()
	mutexLocker.Lock()
	for i := 0; i < 1000; i++ {
		res++
	}
}

func add2() {
	defer wg.Done()
	defer mutexLocker.Unlock()
	mutexLocker.Lock()
	for i := 0; i < 1000; i++ {
		res += 2
	}
}
