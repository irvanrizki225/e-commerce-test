package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID       	uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     	string         `json:"name" gorm:"size:100;not null;unique"`
	Price    	int            `json:"price" gorm:"not null"`
	Stock    	int            `json:"stock" gorm:"not null"`
	CreatedAt 	time.Time      `json:"created_at"`
	UpdatedAt 	time.Time      `json:"updated_at"`
	DeletedAt 	*gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Card struct {
	ID       	uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID	uint           `json:"product_id" gorm:"not null;unique;index;"`
	UserID   	uint           `json:"user_id" gorm:"not null;unique;index;"`
	Quantity 	int            `json:"quantity" gorm:"not null"`
	CreatedAt 	time.Time      `json:"created_at"`
	UpdatedAt 	time.Time      `json:"updated_at"`
	DeletedAt 	*gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Order struct {
	ID       		uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID   		uint           `json:"user_id" gorm:"not null;unique;index;"`
	PaymentMethod 	string         `json:"payment_method" gorm:"size:100;not null;unique"`
	CreatedAt 		time.Time      `json:"created_at"`
	UpdatedAt 		time.Time      `json:"updated_at"`
	DeletedAt 		*gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type OrderDetail struct {
	ID       	uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID  	uint           `json:"order_id" gorm:"not null;unique;index;"`
	ProductID	uint           `json:"product_id" gorm:"not null;unique;index;"`
	CreatedAt 	time.Time      `json:"created_at"`
	UpdatedAt 	time.Time      `json:"updated_at"`
	DeletedAt 	*gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func MigrateProduct(db *gorm.DB) {
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Card{})
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&OrderDetail{})

	fmt.Println("Successfully migrated Product!")
}