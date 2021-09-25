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
	for _, c := range cs {
		c.Name = "Coach:" + c.Name
	}
	return cs
}

func (cp *coachPresenter) ResponseCoach(c *model.Coach) *model.Coach {
	c.Name = "Coach: " + c.Name
	return c
}
