package main

import (
	"fmt"
	"homework/dao"
	"homework/server"
	"log"
)

func init() {
	// 先建立对数据库的连接
	dao.Init()
	//先进行登录才能进行下一步
	for true {
		if server.Login() != nil {
			continue
		}
		break
	}
}

func main() {
	userInput := 0

	for true {
		fmt.Println("1:增加产品\n2:增加用户\n3:等待功能更新")
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			log.Println("user input error")
			fmt.Println("输入读取出错，请重试")
			continue
		}
		switch {
		case userInput == 1:
			log.Println("Add Product")
			if err = server.AddProduct(); err != nil {
				log.Println("Unknown failure in adding product information")
				fmt.Println("添加产品信息出现未知失败，将回到主菜单")
				continue
			} else {
				log.Println("Successfully added product information")
				fmt.Println("添加产品信息成功，将回到主菜单！")
				continue
			}
		case userInput == 2:
			log.Println("Add User")
			if server.AddUser() != "" {
				log.Println("Failed to add user")
				fmt.Println("添加出错啦，将回到主菜单")
				continue
			} else {
				log.Println("Successfully added user")
				fmt.Println("增加用户成功")
				continue
			}
		default:
			fmt.Println("on，这个东西我不认识哦~")
		}
	}
}
