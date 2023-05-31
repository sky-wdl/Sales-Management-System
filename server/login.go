package server

import (
	"fmt"
	"homework/dao"
	"homework/model"
	"log"
)

var loginUser = new(model.User)

func Login() {
	var err error
	for true {
		fmt.Print("请输入账号：")
		if _, err = fmt.Scan(&loginUser); err != nil {
			log.Println("读入账号时出错...")
			fmt.Println("账号读取时发生未知错误,请重试")
			continue
		}
		break
	}
	for true {
		fmt.Println("请输入密码：")
		if _, err = fmt.Scan(&loginUser.Password); err != nil {
			log.Println("读取密码错误。。")
			fmt.Println("密码读入时候发生未知错误，请重试..")
			continue
		}
		break
	}
	if dao.Login(loginUser) != nil {
		log.Println("登录出错啦")
		fmt.Println("登录出错啦")
	} else {
		fmt.Println("登录成功")
	}
}
