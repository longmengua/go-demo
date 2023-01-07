package db

import "gorm.io/gorm"

type Struct struct {
	Host     string
	Port     string
	User     string
	Pwd      string
	DbName   string
	Instance *gorm.DB
}
