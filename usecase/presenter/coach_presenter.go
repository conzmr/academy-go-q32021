package presenter

import "github.com/conzmr/academy-go-q32021/domain/model"

type CoachPresenter interface {
	ResponseCoaches(c []*model.Coach) []*model.Coach
	ResponseCoach(c *model.Coach) *model.Coach
}
