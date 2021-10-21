package interactor

import (
	"testing"

	"github.com/conzmr/academy-go-q32021/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var coaches = []*model.Coach{
	{
		Name: "ex1",
		Id:   1,
	},
	{
		Name: "ex2",
		Id:   2,
	},
}

type MockCoachRepo struct {
	mock.Mock
}

type MockCoachPresenter struct {
	mock.Mock
}

func (mr MockCoachRepo) FindAll(cs []*model.Coach) ([]*model.Coach, error) {
	args := mr.Called()
	return args.Get(0).([]*model.Coach), args.Error(1)
}

func (mr MockCoachRepo) FindById(id string) (*model.Coach, error) {
	args := mr.Called()
	return args.Get(0).(*model.Coach), args.Error(1)
}

func (mp MockCoachPresenter) ResponseCoaches(c []*model.Coach) []*model.Coach {
	args := mp.Called()
	return args.Get(0).([]*model.Coach)
}

func (mp MockCoachPresenter) ResponseCoach(c *model.Coach) *model.Coach {
	args := mp.Called()
	return args.Get(0).(*model.Coach)
}

func TestFindCoaches(t *testing.T) {
	testCases := []struct {
		name     string
		res      []*model.Coach
		hasError bool
		error    error
		id       string
	}{
		{
			name:     "Return coaches list",
			res:      coaches,
			hasError: false,
			error:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCoachRepo := MockCoachRepo{}
			mockCoachPresenter := MockCoachPresenter{}
			mockCoachRepo.On("FindAll").Return(tc.res, tc.error)
			mockCoachPresenter.On("ResponseCoaches").Return(tc.res)
			ci := NewCoachInteractor(mockCoachRepo, mockCoachPresenter)
			cs := []*model.Coach{}
			coaches, err := ci.Get(cs)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.EqualValues(t, tc.res, coaches)
			}
		})
	}
}

func TestFindCoach(t *testing.T) {
	testCases := []struct {
		name     string
		res      *model.Coach
		coachId  string
		hasError bool
		error    error
		id       string
	}{
		{
			name:     "Return specific coach if found",
			res:      coaches[0],
			coachId:  "0",
			hasError: false,
			error:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCoachRepo := MockCoachRepo{}
			mockCoachPresenter := MockCoachPresenter{}
			mockCoachRepo.On("FindById").Return(tc.res, tc.error)
			mockCoachPresenter.On("ResponseCoach").Return(tc.res)
			ci := NewCoachInteractor(mockCoachRepo, mockCoachPresenter)
			coach, err := ci.GetById(tc.coachId)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.EqualValues(t, tc.res, coach)
			}
		})
	}
}
