package model

import (
	"strconv"
)

type Exercise struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (e Exercise) ToCSV() []string {
	return []string{strconv.Itoa(e.Id), e.Name}
}
