package models

import (
	"time"
)

type User struct {
	ID			int			`json:"id" gorm:"primaryKey"`
	Name		string		`json:"name" gorm:"type:varchar(255)"`		
	CreatedAt	time.Time	`json:"created_at"`
	Email		string		`json:"email" gorm:"type:varchar(255)"`
}

