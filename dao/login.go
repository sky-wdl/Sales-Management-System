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
			log.Println(err)
			return err
		}
		return nil
	})
	return err
}
