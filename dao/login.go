package dao

import (
	"gorm.io/gorm"
	"homework/model"
	"log"
)

func Login(loginUser *model.User) (err error) {
	err = userDb.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Table("users").
			Where("userName = ? AND passWord = ?", loginUser.UserName, loginUser.Password).
			First(&loginUser).Error
		if err != nil {
			log.Println("查询账号时出错，不知道是账号还是密码错了")
			return err
		}
		return nil
	})
	return err
}
