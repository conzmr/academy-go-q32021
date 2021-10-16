package model

type Coach struct {
	Id   int    `csv:"id" json:"id"`     // Id of the coach
	Name string `csv:"Name" json:"name"` // Name of the coach
}

func (Coach) TableName() string { return "coaches" }
