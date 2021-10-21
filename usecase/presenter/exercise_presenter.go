package presenter

import "github.com/conzmr/academy-go-q32021/domain/model"

type ExercisePresenter interface {
	ResponseExercises(c []*model.Exercise) []*model.Exercise
}
