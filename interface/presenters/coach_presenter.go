package presenter

import (
	"academy-go-q32021/academy-go-q32021/domain/model"
)

type coachPresenter struct {
}

type CoachPresenter interface {
	ResponseCoaches(us []*model.Coach) []*model.Coach
}

func NewUserPresenter() CoachPresenter {
	return &coachPresenter{}
}

func (cp *coachPresenter) ResponseCoaches(cs []*model.Coach) []*model.Coach {
	for _, c := range cs {
		c.Name = "Coach:" + c.Name
	}
	return cs
}
