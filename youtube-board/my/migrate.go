package my

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() {
	dsn := "host=localhost dbname=postgres sslmode=disable"
	v2Db, er := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if er != nil {
		fmt.Println(er)
		return
	}
	db, er := v2Db.DB()
	if er != nil {
		fmt.Println(er)
		return
	}
	defer db.Close()
	v2Db.AutoMigrate(&User{}, &Group{}, &Post{}, &Comment{})
}
