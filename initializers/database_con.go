package initializers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./rants.db"))

	if err != nil {
		log.Fatal("Failed to open database")
	}

	// db.Create(&User{
	// 	Name: "Melusi",
	// 	Surname: "Gumbi",
	// 	Age: 22,
	// })

	// var users[]

}
