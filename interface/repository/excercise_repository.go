package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/conzmr/academy-go-q32021/domain/model"
)

type ExerciseResponse struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []*model.Exercise `json:"results"`
}

type exerciseRepository struct {
}

type ExerciseRepository interface {
	FindAll(c []*model.Exercise) ([]*model.Exercise, error)
	// Save(c []*model.Exercise) ([]*model.Exercise, error)
}

func NewExerciseRepository() ExerciseRepository {
	return &exerciseRepository{}
}

func (er *exerciseRepository) FindAll(c []*model.Exercise) ([]*model.Exercise, error) {

	url := fmt.Sprintf("https://wger.de/api/v2/exercise/?format=json&limit=%d", 10)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyRaw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var exercisesRes ExerciseResponse

	if err = json.Unmarshal(bodyRaw, &exercisesRes); err != nil {
		return nil, err
	}
	exercises := exercisesRes.Results
	return exercises, nil
}
