package model

type Coach struct {
	Id   int    `csv:"id"`
	Name string `csv:"Name"`
}

func (Coach) TableName() string { return "coaches" }
