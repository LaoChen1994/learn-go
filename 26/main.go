package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// delay := time.NewTimer(2 * time.Second)
	// fmt.Println("倒计时开始")
	// defer delay.Stop()

	// delaySeconds, ok := <-delay.C

	// if ok {
	// 	fmt.Println("倒计时结束", delaySeconds, "s")
	// }

	// ticker := time.NewTicker(2 * time.Second)
	// fmt.Println("每过2s常态化触发开始")
	// defer ticker.Stop()

	// end := 0

	// for val := range ticker.C {
	// 	fmt.Println("有过了两秒啦~", val)

	// 	end++

	// 	if end > 10 {
	// 		return
	// 	}
	// }

	syncTask()
}

type process struct {
	current int
	total   int
}

var wg sync.WaitGroup

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
