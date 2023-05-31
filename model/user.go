package model

type User struct {
	ID       int8   `gorm:"column:id"`
	UserName string `gorm:"column:userName"`
	Password string `gorm:"column:password"`
	Level    int8   `gorm:"column:level"`
	//Time     time.Time `gorm:"column:time"`
}

func (User) TabName() string {
	return "users"
}
