package presenter

import (
	"github.com/conzmr/academy-go-q32021/domain/model"
)

type exercisePresenter struct {
}

type ExercisePresenter interface {
	ResponseExercises(cs []*model.Exercise) []*model.Exercise
}

// Returns an ExercisePresenter
func NewExercisePresenter() ExercisePresenter {
	return &exercisePresenter{}
}

// Handles exercises data before passing it to view
func (ep *exercisePresenter) ResponseExercises(ex []*model.Exercise) []*model.Exercise {
	return ex
}
