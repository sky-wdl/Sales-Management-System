package dao

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"homework/model"
	"log"
)

func AddProduct(p *model.Product) (err error) {
	//储存查询到的内容
	var productTemp = new(model.Product)

	//根据PK码先查询库里是否已经有这个商品，有的话就报错，要求手动针对库存进行更新，没有这个商品就顺序执行下去
	err = db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Table(p.TabName()).
			Where("PK = ?", p.PK).
			First(&productTemp).Error; err == nil {
			log.Println("未找到商品")
		} else {
			return
		}
		return nil
	})
	if err != nil {
		fmt.Println("商品已存在...")
		return err
	}

	// 开启事务
	err = db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Table(p.TabName()).
			Clauses(clause.OnConflict{DoNothing: true}).
			Create(p).Error; err != nil {
			panic(err)
			return err
		} else {
			return nil
		}
	})
	return err
}
