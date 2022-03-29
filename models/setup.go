package models
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Recent_apps string `json:"recent_apps"`
	Favourite_apps string `json:"favourite_apps"`
}
var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{})
	DB = database

}
