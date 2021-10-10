package presenter

import (
	"github.com/conzmr/academy-go-q32021/domain/model"
)

type coachPresenter struct {
}

type CoachPresenter interface {
	ResponseCoaches(cs []*model.Coach) []*model.Coach
	ResponseCoach(c *model.Coach) *model.Coach
}

func NewCoachPresenter() CoachPresenter {
	return &coachPresenter{}
}

func (cp *coachPresenter) ResponseCoaches(cs []*model.Coach) []*model.Coach {
	return cs
}

func (cp *coachPresenter) ResponseCoach(c *model.Coach) *model.Coach {
	return c
}
