package controller

import (
	"net/http"

	"github.com/conzmr/academy-go-q32021/domain/model"
	"github.com/conzmr/academy-go-q32021/usecase/interactor"
)

type exerciseController struct {
	exerciseInteractor interactor.ExerciseInteractor
}

type ExerciseController interface {
	GetExercises(c Context) error
	SyncExercises(c Context) error
}

func NewExerciseController(e interactor.ExerciseInteractor) ExerciseController {
	return &exerciseController{e}
}

func (ec *exerciseController) GetExercises(c Context) error {
	var e []*model.Exercise

	e, err := ec.exerciseInteractor.Get(e)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, e)
}

func (ec *exerciseController) SyncExercises(c Context) error {
	var e []*model.Exercise

	e, err := ec.exerciseInteractor.Get(e)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, e)
}
