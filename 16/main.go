package main

import (
	"fmt"
	"reflect"
)

type FileDest interface {
	chooseDest(url string, name string) string
}

type Save interface {
	saveFile()
}

type FileDownload interface {
	FileDest
	Save
}

type File struct {
	path string
	name string
}

func (f *File) chooseDest(url string, name string) string {
	f.path = url
	f.name = name
	return url
}

func (f *File) saveFile() {
	fmt.Println(f.name, "存储成功, 地址", f.path)
}

func dataOutput(data interface{}) {
	dataVal, stringOk := data.(string)

	if stringOk {
		fmt.Println(dataVal)
	}
}

func getInitLoader() *FileDownload {
	var downloader *FileDownload = nil

	return downloader
}

func main() {
	var downloader FileDownload

	downloader = new(File)
	downloader.chooseDest("/user/download/1.txt", "1.txt")
	downloader.saveFile()

	fmt.Println(reflect.TypeOf(downloader))

	var a = [5]int{1, 2, 3, 4, 5}
	var b = [5]int{1, 2, 3, 4, 5}

	fmt.Println(a == b)

	fmt.Println(getInitLoader())
	fmt.Println(getInitLoader() == nil)
}
