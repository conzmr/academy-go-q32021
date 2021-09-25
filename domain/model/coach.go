package model

type Coach struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func (Coach) TableName() string { return "coaches" }
