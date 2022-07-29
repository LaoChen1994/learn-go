package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	var filePath = "E:/Learn/learn-go/21/1.txt"

	createFile(filePath)

	file, err := os.OpenFile(filePath, os.O_APPEND, os.ModePerm)

	content, error := ioutil.ReadFile(filePath)

	if error != nil {
		panic(error)
	}

	val := string(content)

	fmt.Println(val)

	if err != nil {
		panic(err)
	}

	str := strings.Join(args, " ") + "\n"
	file.Write([]byte(str))
	file.Close()
}

func createFile(filePath string) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(filePath)

			if err != nil {
				fmt.Println("创建文件失败, ", err)
				panic(err)
			}

			file.Close()
			return
		} else {
			panic(err)
		}
	}

	if fileInfo.IsDir() {
		panic("已存在文件为文件夹，请确认路径")
	}
}
