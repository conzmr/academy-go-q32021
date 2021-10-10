package repository

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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

	csvfile, err := os.OpenFile("./datastore/exercises.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	defer writer.Flush()

	for _, record := range exercises {
		err := writer.Write(record.ToCSV())
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
	}

	return exercises, nil
}
