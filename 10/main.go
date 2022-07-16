package main

import (
	"fmt"
	"time"
)

var precent int = 0
var isEnd = false

func main() {
	func(text string) {
		fmt.Println(text)
	}("立即使用")

	stringLoop := func(text string) {
		fmt.Println(text)
	}

	stringLoop("传参使用")

	go download("", func() {
		fmt.Println("下载成功")
		isEnd = true
	}, func() {
		fmt.Println("下载失败，当前进度", precent, "%")
		isEnd = true
	})

	for {
		if precent == 10 {
			isEnd = true
			time.Sleep(20 * time.Second)
		}

		if isEnd {
			break
		} else {
			time.Sleep(1 * time.Second)
			fmt.Println("当前进度", precent, "%")
		}
	}

}

func download(url string, onSuccess func(), onFail func()) {
	for {
		if isEnd {
			if precent != 100 {
				onFail()
			}

			break
		}

		time.Sleep(1 * time.Second)
		precent += 1

		if precent >= 100 {
			onSuccess()
