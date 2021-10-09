package model

type Coach struct {
	Id   int    `csv:"id" json:"id"`
	Name string `csv:"Name" json:"name"`
}

func (Coach) TableName() string { return "coaches" }
