package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/conzmr/academy-go-q32021/domain/model"
)

const (
	endpoint                = "https://api.openweathermap.org/data/2.5"
	pathFormatWeatherByCity = "/weather?q=%s&appid=%s&units=metric"
)

type ExerciseResponse struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []*model.Exercise `json:"exercises"`
}

type exerciseRepository struct {
}

type ExerciseRepository interface {
	FindAll(c []*model.Exercise, limit int) ([]*model.Exercise, error)
	// Save(c []*model.Exercise) ([]*model.Exercise, error)
}

func NewExerciseRepository() ExerciseRepository {
	return &exerciseRepository{}
}

func (er *exerciseRepository) FindAll(c []*model.Exercise, limit int) ([]*model.Exercise, error) {

	url := fmt.Sprintf("https://wger.de/api/v2/exercise/?format=json&limit=%d", limit)

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("openweather.GetWeatherByCity failed http GET: %s", err)
	}
	defer res.Body.Close()

	// read the response body and encode it into the respose struct
	bodyRaw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("openweather.GetWeatherByCity failed reading body: %s", err)
	}

	var exercisesRes ExerciseResponse
	exercises := []*model.Exercise{}

	if err = json.Unmarshal(bodyRaw, &exercisesRes); err != nil {
		return nil, fmt.Errorf("openweather.GetWeatherByCity failed encoding body: %s", err)
	}

	exercises = exercisesRes.Results
	return exercises, nil
}
