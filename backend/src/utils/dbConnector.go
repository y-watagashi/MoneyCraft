package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
	"reflect"
)

func ConnectToDB() *gorm.DB{
	dsn := "root:asdf@tcp(db:3306)/main?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println(reflect.TypeOf(db))
	return db
}