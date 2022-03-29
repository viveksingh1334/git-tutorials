package models

type User struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Recent_apps string `json:"recent_apps"`
	Favourite_apps string `json:"favourite_apps"`
}