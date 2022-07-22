package main

import "fmt"

type LearnCode interface {
	LearnBasic() string
	StartCoding(code string) string
	FindWork(company string) string
}

type StudyGo struct {
	language string
}

func (coder StudyGo) LearnBasic() string {
	return "学习" + coder.language
}

func (coder StudyGo) StartCoding(code string) string {
	return "使用" + coder.language + "写" + code
}

func (coder StudyGo) FindWork(company string) string {
	return "去" + company + "写" + coder.language
}

func main() {
	var goer LearnCode

	goer = new(StudyGo)

	fmt.Println(goer.LearnBasic())
	fmt.Println(goer.StartCoding("Hello World"))
	fmt.Println(goer.FindWork("bytedance"))
}
