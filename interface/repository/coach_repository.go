package repository

import (
	"log"
	"os"

	"github.com/conzmr/academy-go-q32021/domain/model"

	"github.com/gocarina/gocsv"
)

type coachRepository struct {
}

type CoachRepository interface {
	FindAll(c []*model.Coach) ([]*model.Coach, error)
}

func NewCoachRepository() CoachRepository {
	return &coachRepository{}
}

func (cr *coachRepository) FindAll(c []*model.Coach) ([]*model.Coach, error) {
	coachFile, err := os.OpenFile("./datastore/coach.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer coachFile.Close()
	coaches := []*model.Coach{}

	if err := gocsv.UnmarshalFile(coachFile, &coaches); err != nil {
		log.Fatal(err)
	}
	return coaches, nil

	if err != nil {
		return nil, err
	}

	return c, nil
}
