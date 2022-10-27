package models

import (
	"time"
)

type Book struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     *string   `json:"title" gorm:"not null"`
	Isbn      *string   `json:"isbn" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm: "autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm: "autoCreateTime"`
}
