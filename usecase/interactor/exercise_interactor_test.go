package interactor

import (
	"testing"

	"github.com/conzmr/academy-go-q32021/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var exercises = []*model.Exercise{
	{
		Name: "ex1",
		Id:   1,
	},
	{
		Name: "ex2",
		Id:   2,
	},
	{
		Name: "ex3",
		Id:   3,
	},
}

type ExerciseMockRepo struct {
	mock.Mock
}

type ExerciseMockPresenter struct {
	mock.Mock
}

func (mr ExerciseMockRepo) FindAll(exs []*model.Exercise, typeId string, items int, itemsPerWorkers int) ([]*model.Exercise, error) {
	args := mr.Called()
	return args.Get(0).([]*model.Exercise), args.Error(1)
}

func (mr ExerciseMockRepo) Sync(exs []*model.Exercise) ([]*model.Exercise, error) {
	args := mr.Called()
	return args.Get(0).([]*model.Exercise), args.Error(1)
}

func (mp ExerciseMockPresenter) ResponseExercises(c []*model.Exercise) []*model.Exercise {
	args := mp.Called()
	return args.Get(0).([]*model.Exercise)
}

func TestFindExercises(t *testing.T) {
	testCases := []struct {
		name            string
		ex              []*model.Exercise
		hasError        bool
		error           error
		id              string
		typeId          string
		items           int
		itemsPerWorkers int
	}{
		{
			name:            "Return exercises list",
			ex:              exercises,
			hasError:        false,
			error:           nil,
			typeId:          "Odd",
			items:           3,
			itemsPerWorkers: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			exerciseMockRepo := ExerciseMockRepo{}
			exerciseMockPresenter := ExerciseMockPresenter{}
			exerciseMockRepo.On("FindAll").Return(tc.ex, tc.error)
			exerciseMockPresenter.On("ResponseExercises").Return(tc.ex)
			uei := NewExerciseInteractor(exerciseMockRepo, exerciseMockPresenter)
			exs := []*model.Exercise{}
			exercises, err := uei.Get(exs, tc.typeId, tc.items, tc.itemsPerWorkers)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.EqualValues(t, tc.ex, exercises)
			}
		})
	}
}

func TestSyncExercises(t *testing.T) {
	testCases := []struct {
		name     string
		ex       []*model.Exercise
		hasError bool
		error    error
		id       string
	}{
		{
			name:     "Return syn exercises",
			ex:       exercises,
			hasError: false,
			error:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockExerciseRepo := ExerciseMockRepo{}
			mockExercisePresenter := ExerciseMockPresenter{}
			mockExerciseRepo.On("Sync").Return(tc.ex, tc.error)
			mockExercisePresenter.On("ResponseExercises").Return(tc.ex)
			uei := NewExerciseInteractor(mockExerciseRepo, mockExercisePresenter)
			exs := []*model.Exercise{}
			exercises, err := uei.Sync(exs)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.EqualValues(t, tc.ex, exercises)
			}
		})
	}
}
