package model

type saledetail struct {
	SerialNumber string
	BarCode      string
}

func (saledetail) TabName() string {
	return "tsaledetail"
}
