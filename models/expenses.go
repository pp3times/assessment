package models

import "gorm.io/gorm"

type Expenses struct {
	ID uint `gorm:"primary key;autoIncrement" json:"id"`

	Title *string `json:"title"`

	Amount *float64 `json:"amount"`

	Note *string `json:"note"`

	Tags []string `json:"tags" gorm:"serializer:json"`
}

func MigrateBooks(db *gorm.DB) error {

	err := db.AutoMigrate(&Expenses{})

	return err

}
