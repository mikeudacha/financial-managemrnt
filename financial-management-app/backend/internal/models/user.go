package models

type User struct {
	ID       int    `gorm:"primary_key"`
	Username string `gorm:"unique not null"`
	Password string `gorm:"not null"`
}
