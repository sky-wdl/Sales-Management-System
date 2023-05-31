package server

import (
	"fmt"
	"homework/model"
	"log"
)

var addUser = new(model.User)

func AddUser() string {
	if loginUser.Level < 3 {
		log.Println("Insufficient account permission level")
		fmt.Println("啊偶，你当前登录的账号等级不足欸，这咋搞啊...")
		return "Insufficient account permission level"
	}

	for true {
		fmt.Print("请输入注册用户名（英文字母）:")
		if _, err := fmt.Scanln(&addUser.UserName); err != nil {
			log.Println("Input error")
			fmt.Println("读取输入的时候出现错误欸，咋搞，要不再来一遍？")
			continue
		}
		break
	}
	for true {
		fmt.Print("输入一下你的账号密码:")
	}
	return ""
}
