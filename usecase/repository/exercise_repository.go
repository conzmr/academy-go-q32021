package repository

import (
	"github.com/conzmr/academy-go-q32021/domain/model"
)

type ExerciseRepository interface {
	FindAll(exs []*model.Exercise, t string, i int, ipw int) ([]*model.Exercise, error)
	Sync(exs []*model.Exercise) ([]*model.Exercise, error)
}
