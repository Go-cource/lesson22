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

	// newUser := User{
	// 	FirstName: "Leha",
	// 	LastName:  "Lehin",
	// 	Email:     "Leha@mail.ru",
	// 	Country:   "Russia",
	// 	Age:       29,
	// }

	// result := db.Create(&newUser)
	// if result.Error != nil {
	// 	fmt.Println("error adding leha", result.Error)
	// 	return
	// }

	// var users []User
	// result := db.Find(&users)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// for i := range users {
	// 	users[i].Age += 1
	// }
	// result = db.Save(&users)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// var user User
	// result := db.First(&user)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	// user.Age = 31
	// result = db.Save(&user)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	var user User
	result := db.First(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	result = db.Delete(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
	} else if result.RowsAffected == 0 {
		fmt.Println("No rows affected")
	} else {
		fmt.Println("Deleted successfully")
	}
	

}
