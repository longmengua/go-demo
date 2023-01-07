package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (s *Struct) Init() {
	// dsn := "host=db.mqkycbwhynxeenavanil.supabase.co user=postgres password=ifHCR0kfZKVqMFfd dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Taipei"
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		s.Host,
		s.User,
		s.Pwd,
		s.DbName,
		s.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect Database")
	}
	s.Instance = db
}
