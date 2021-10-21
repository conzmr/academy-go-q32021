package model

import (
	"strconv"
)

type Exercise struct {
	Id   int    `json:"id" csv:"Id"`     // Id of the exercise
	Name string `json:"name" csv:"Name"` // Name of the exercise
}

// Converts Exercise struct to an array of strings
func (e Exercise) ToCSV() []string {
	return []string{strconv.Itoa(e.Id), e.Name}
}
