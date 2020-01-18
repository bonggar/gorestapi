package model

import (
	"time"
)

//Model : base model that store minimum fields to be a model
type Model struct {
	ID        uint      `gorm:"column:id;primary_key;AUTO_INCREMENT" form:"id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
}
