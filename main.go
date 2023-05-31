package main

import (
	"fmt"
	"homework/dao"
	"homework/server"
	"log"
)

func init() {
	dao.Init()
}

func main() {
	//打印正常运行的提示信息
	log.Println("system starts running")
	userInput := 0
	fmt.Println("1:增加产品\n2:增加用户\n3:登录账户")

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
			if err = server.AddProduct(); err != nil {
				log.Println("Unknown failure in adding product information")
				fmt.Println("添加产品信息出现未知失败")
				continue
			} else {
				log.Println("Successfully added product information")
				fmt.Println("添加产品信息成功！")
				continue
			}
		case userInput == 2:
			fmt.Println("输入的是2")
			if server.AddUser() != "" {
				log.Println("登录出错啦")
			} else {
				log.Println("登录成功")
				fmt.Println("登录成功")
			}
		default:
			fmt.Println("on，这个东西我不认识哦~")
		}
	}
}
