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
	GetCoach(c Context) error
}

func NewCoachController(c interactor.CoachInteractor) CoachController {
	return &coachController{c}
}

func (cc *coachController) GetCoaches(c Context) error {
	var co []*model.Coach

	co, err := cc.coachInteractor.Get(co)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, co)
}

func (cc *coachController) GetCoach(c Context) error {
	var co *model.Coach

	id := c.Param("id")

	co, err := cc.coachInteractor.GetById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, co)
}
