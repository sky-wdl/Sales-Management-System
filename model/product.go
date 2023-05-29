package model

type Product struct {
	PK        int `gorm:"primarykey"`
	Commodity string
	Price     float32
	Stock     int
	Supplier  string
}

func (Product) TabName() string {
	return "tproduct"
}
