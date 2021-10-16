package repository

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/conzmr/academy-go-q32021/domain/model"
)

var file = "./datastore/exercises.csv"

type ExerciseResponse struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []*model.Exercise `json:"results"`
}

type exerciseRepository struct {
}

type ExerciseRepository interface {
	Sync(c []*model.Exercise) ([]*model.Exercise, error)
	FindAll(c []*model.Exercise, typeId string, items int, itemsPerWorkers int) ([]*model.Exercise, error)
}

func NewExerciseRepository() ExerciseRepository {
	return &exerciseRepository{}
}

func (er *exerciseRepository) FindAll(c []*model.Exercise, typeId string, items int, itemsPerWorkers int) ([]*model.Exercise, error) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fcsv := csv.NewReader(f)
	rs := make([]*model.Exercise, 0)
	numWps := int(math.Ceil(float64(items) / float64(itemsPerWorkers)))
	jobs := make(chan []string, numWps)
	res := make(chan *model.Exercise)

	var wg sync.WaitGroup
	worker := func(jobs <-chan []string, results chan<- *model.Exercise) {
		done := 0
		for {
			select {
			case job, ok := <-jobs:
				if !ok || done >= itemsPerWorkers {
					return
				}
				ex, err := parseExercise(job)
				if err != nil || typeId == "even" && ex.Id%2 != 0 {
					break
				}
				done++
				results <- ex
			}
		}
	}

	for w := 0; w < numWps; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, res)
		}()
	}

	go func() {
		for {
			rStr, err := fcsv.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			jobs <- rStr
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		rs = append(rs, r)
	}
	return rs, nil
}

func (er *exerciseRepository) Sync(c []*model.Exercise) ([]*model.Exercise, error) {

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

	csvfile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
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

func parseExercise(data []string) (*model.Exercise, error) {
	id, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, err
	}
	return &model.Exercise{
		Id:   id,
		Name: data[1],
	}, nil
}
