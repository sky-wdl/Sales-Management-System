package model

type Product struct {
	PK        int     `gorm:"primarykey"`
	Commodity string  `gorm:"column:commodity"`
	Price     float32 `gorm:"column:price"`
	Stock     int     `gorm:"column:stock"`
	Supplier  string  `gorm:"column:supplier"`
}

func (Product) TabName() string {
	return "tproduct"
}
