package repository

import (
	"academy-go-q32021/academy-go-q32021/domain/model"

	"github.com/jinzhu/gorm"
)

type coachRepository struct {
	db *gorm.DB
}

type CoachRepository interface {
	FindAll(u []*model.Coach) ([]*model.Coach, error)
}

func NewCoachRepository(db *gorm.DB) CoachRepository {
	return &coachRepository{db}
}

func (cr *coachRepository) FindAll(c []*model.Coach) ([]*model.Coach, error) {
	err := cr.db.Find(&c).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}
