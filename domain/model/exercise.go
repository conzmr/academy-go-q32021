package model

import (
	"strconv"
)

type Exercise struct {
	Id   int    `json:"id" csv:"Id"`
	Name string `json:"name" csv:"Name"`
}

func (e Exercise) ToCSV() []string {
	return []string{strconv.Itoa(e.Id), e.Name}
}
