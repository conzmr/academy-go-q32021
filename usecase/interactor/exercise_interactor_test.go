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

type MockRepo struct {
	mock.Mock
}

type MockPresenter struct {
	mock.Mock
}

func (mr MockRepo) FindAll(exs []*model.Exercise) ([]*model.Exercise, error) {
	args := mr.Called()
	return args.Get(0).([]*model.Exercise), args.Error(1)
}

func (mp MockPresenter) ResponseExercises(c []*model.Exercise) []*model.Exercise {
	args := mp.Called()
	return args.Get(0).([]*model.Exercise)
}

func TestFindExercises(t *testing.T) {
	testCases := []struct {
		name     string
		ex       []*model.Exercise
		hasError bool
		error    error
		id       string
	}{
		{
			name:     "Return exercises list",
			ex:       exercises,
			hasError: false,
			error:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := MockRepo{}
			mockPresenter := MockPresenter{}
			mockRepo.On("FindAll").Return(tc.ex, tc.error)
			mockPresenter.On("ResponseExercises").Return(tc.ex)
			uei := NewExerciseInteractor(mockRepo, mockPresenter)
			exs := []*model.Exercise{}
			exercises, err := uei.Get(exs)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.EqualValues(t, tc.ex, exercises)
			}
		})
	}
}
