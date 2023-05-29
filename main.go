package main

import (
	"fmt"
	"homework/dao"
	"log"
)

func init() {
	dao.Init()
}

func main() {
	//打印正常运行的提示信息
	log.Println("system starts running")
	userInput := 0
	fmt.Println("1:增加产品\n2:增加用户")

	//打印选择选项
	for true {
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			log.Println("user input error")
			fmt.Println("输入读取出错，请重试")
			continue
		}
		switch {
		case userInput == 1:
			fmt.Println("输入的是1")
		default:
			fmt.Println("on，这个东西我不认识哦~")
		}
	}
}
