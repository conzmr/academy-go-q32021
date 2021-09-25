package presenter

import "academy-go-q32021/academy-go-q32021/domain/model"

type CoachPresenter interface {
	ResponseUsers(u []*model.Coach) []*model.Coach
}
