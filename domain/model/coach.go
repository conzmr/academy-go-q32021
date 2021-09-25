package model

import "time"

type Coach struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (Coach) TableName() string { return "coaches" }
