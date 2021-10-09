package model

type Exercise struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (Exercise) TableName() string { return "exercises" }
