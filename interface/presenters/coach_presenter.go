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

// Returns an CoachPresenter
func NewCoachPresenter() CoachPresenter {
	return &coachPresenter{}
}

// Handles coach data before passing it to view
func (cp *coachPresenter) ResponseCoaches(cs []*model.Coach) []*model.Coach {
	return cs
}

// Handles coaches data before passing it to view
func (cp *coachPresenter) ResponseCoach(c *model.Coach) *model.Coach {
	return c
}
