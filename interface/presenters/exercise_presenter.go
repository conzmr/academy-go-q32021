package presenter

import (
	"github.com/conzmr/academy-go-q32021/domain/model"
)

type exercisePresenter struct {
}

type ExercisePresenter interface {
	ResponseExercises(cs []*model.Exercise) []*model.Exercise
}

func NewExercisePresenter() ExercisePresenter {
	return &exercisePresenter{}
}

func (ep *exercisePresenter) ResponseExercises(ex []*model.Exercise) []*model.Exercise {
	return ex
}
