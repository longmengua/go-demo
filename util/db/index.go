package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (s *Struct) Init(
	host string,
	port string,
	user string,
	pwd string,
) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/TPDB?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, port)
	instance, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect MySQL database")
	}
	s.Instance = instance
}
