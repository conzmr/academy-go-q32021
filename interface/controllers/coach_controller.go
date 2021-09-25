package controller

import (
	"net/http"

	"github.com/conzmr/academy-go-q32021/domain/model"
	"github.com/conzmr/academy-go-q32021/usecase/interactor"
)

type coachController struct {
	coachInteractor interactor.CoachInteractor
}

type CoachController interface {
	GetCoaches(c Context) error
}

func NewCoachController(c interactor.CoachInteractor) CoachController {
	return &coachController{c}
}

func (cc *coachController) GetCoaches(ctx Context) error {
	var c []*model.Coach

	c, err := cc.coachInteractor.Get(c)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, c)
}
