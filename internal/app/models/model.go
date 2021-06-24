package models

import "time"

type BaseModel struct {
	ID uint64

	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}
