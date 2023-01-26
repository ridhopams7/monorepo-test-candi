package domain

import (
	"time"
)

// Brand model
type Brand struct {
	ID        string    `gorm:"column:id;type:varchar(255);primary_key" json:"id"`
	Title     string    `gorm:"column:title;type:varchar(255)" json:"title"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName return table name of Brand model
func (Brand) TableName() string {
	return "brands"
}
