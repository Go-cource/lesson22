package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"uniqueIndex"`
	LastName  string `gorm:"uniqueIndex"`
	Email     string `gorm:"not null"`
	Country   string `gorm:"not null"`
	Age       int    `gorm:"not null;size:3"`
}

func main() {
	pgConnect := "host=localhost user=postgres password=yd4v dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(pgConnect), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}

}
