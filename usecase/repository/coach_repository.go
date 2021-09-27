package repository

import (
	"github.com/conzmr/academy-go-q32021/domain/model"
)

type CoachRepository interface {
	FindAll(c []*model.Coach) ([]*model.Coach, error)
	FindById(id string) (*model.Coach, error)
}
