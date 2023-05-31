package server

import (
	"fmt"
	"homework/dao"
	"homework/model"
	"log"
)

var loginUser = new(model.User)

func Login() (err error) {
	for true {
		fmt.Print("请输入账号：")
		if _, err = fmt.Scanln(&loginUser.UserName); err != nil {
			log.Println("Unknown error in input")
			fmt.Println("账号读取时发生未知错误,请重试")
			continue
		}
		break
	}
	for true {
		fmt.Print("请输入密码：")
		if _, err = fmt.Scanln(&loginUser.Password); err != nil {
			log.Println("读取密码错误。。")
			fmt.Println("密码读入时候发生未知错误，请重试..")
			continue
		}
		break
	}

	if err = dao.Login(loginUser); err != nil {
		log.Println("Login failed")
		fmt.Println("请重试..")
		return err
	} else {
		log.Println("Login succeeded")
		fmt.Println("登录成功")
		return nil
	}
}
