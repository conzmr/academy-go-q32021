package model

type Exercise struct {
	Id          int
	Name        string
	Description string
}

func (Exercise) TableName() string { return "exercises" }
