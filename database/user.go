package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `json:"id" gorm:"primaryKey"`
	UserName string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}
