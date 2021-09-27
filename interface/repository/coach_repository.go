package repository

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/conzmr/academy-go-q32021/domain/model"

	"github.com/gocarina/gocsv"
)

type coachRepository struct {
}

type CoachRepository interface {
	FindAll(c []*model.Coach) ([]*model.Coach, error)
	FindById(id string) (*model.Coach, error)
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

func (cr *coachRepository) FindById(id string) (*model.Coach, error) {
	fmt.Println("Message hee")
	coachFile, err := os.Open("./datastore/coach.csv")
	if err != nil {
		return nil, err
	}
	defer coachFile.Close()
	coach := []*model.Coach{}
	reader := csv.NewReader(coachFile)
	var headers []string
	var csvString string
	i := 0
	for {
		record, err := reader.Read()
		if i == 0 {
			headers = record
			i++
			csvString = strings.Join(headers, ",") + "\n"
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] == id {
			csvString += strings.Join(record, ",")
			break
		}
	}
	if err := gocsv.UnmarshalString(csvString, &coach); err != nil {
		log.Fatal(err)
	}
	return coach[0], nil
}
