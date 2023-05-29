package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 定义一个数据库对象
var db *gorm.DB
var userDb *gorm.DB

func Init() {
	//连接销售数据库sale
	dbDsn := "root:12345678@tcp(127.0.0.1:3306)/sale?charset=utf8mb4&parseTime=True&loc=Local"
	dbTemp, err := gorm.Open(mysql.Open(dbDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		log.Println("Successfully connected to sale database!")
		db = dbTemp
	}

	// 连接账户数据库sale_user
	userDbDsn := "root:12345678@tcp(127.0.0.1:3306)/sale_user?charset=utf8mb4&parseTime=True&loc=Local"
	userDbTemp, err := gorm.Open(mysql.Open(userDbDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		log.Printf("Successfully connected to sale_user database!")
		userDb = userDbTemp
	}
}
