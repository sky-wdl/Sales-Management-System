package server

import (
	"errors"
	"fmt"
	"homework/dao"
	"homework/model"
)

func AddProduct() (err error) {
	var p = new(model.Product)
	for true {
		fmt.Print("请输入商品的条码：")
		if _, err = fmt.Scan(&p.PK); err != nil {
			fmt.Println("读取输入出错，请重试")
			continue
		}
		break
	}
	for true {
		fmt.Print("请输入商品的名称：")
		if _, err = fmt.Scan(&p.Commodity); err != nil {
			fmt.Println("读取输入出错，请重试")
			continue
		}
		break
	}
	for true {
		fmt.Print("请输入产品的价格：")
		if _, err = fmt.Scan(&p.Price); err != nil {
			fmt.Println("读取输入出错，请重试")
			continue
		}
		break
	}
	for true {
		fmt.Print("请输入产品的库存数：")
		if _, err = fmt.Scan(&p.Stock); err != nil {
			fmt.Println("读取输入出错，请重试")
			continue
		}
		break
	}
	for true {
		fmt.Print("请输入产品的供应商：")
		if _, err = fmt.Scan(&p.Supplier); err != nil {
			fmt.Println("读取输入出错，请重试")
			continue
		}
		break
	}

	if dao.AddProduct(p) == nil {
		err = errors.New("item already exists")
		return err
	} else {
		return nil
	}
}
