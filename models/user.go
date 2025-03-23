package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       	uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Username    string         `json:"username" gorm:"type:varchar(100);uniqueIndex"`
	Email    	string         `json:"email" gorm:"type:varchar(100);uniqueIndex"`
	Password 	string         `json:"password" gorm:"size:255;not null"`
	Token    	string         `json:"token"`
	CreatedAt 	time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt 	time.Time      `json:"updated_at"`
	DeletedAt 	*gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func MigrateUser(db *gorm.DB) {
	db.AutoMigrate(&User{})

	fmt.Println("Successfully migrated User!")
}