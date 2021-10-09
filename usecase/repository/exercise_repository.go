package repository

import (
	"github.com/conzmr/academy-go-q32021/domain/model"
)

type ExerciseRepository interface {
	FindAll(exs []*model.Exercise) ([]*model.Exercise, error)
}
