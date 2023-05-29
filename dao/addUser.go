package dao

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"homework/model"
)

// Adduser 增加用户
func Adduser(name, passwd string) (err error) {

	// 拼凑用户账号信息
	userTemp := model.User{
		UserName: name,
		Password: passwd,
		Level:    1, //默认是1
	}

	err = userDb.Transaction(func(tx *gorm.DB) (err error) {
		if err := userDb.Table("users").
			Clauses(clause.OnConflict{DoNothing: true}).
			Create(&userTemp).Error; err != nil {
			return err
		} else {
			return nil
		}
	})

	return err
}
