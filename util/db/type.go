package db

import "gorm.io/gorm"

type Struct struct {
	Instance *gorm.DB
}
