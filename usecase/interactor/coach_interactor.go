package interactor

import (
	"github.com/conzmr/academy-go-q32021/domain/model"

	"github.com/conzmr/academy-go-q32021/usecase/presenter"
	"github.com/conzmr/academy-go-q32021/usecase/repository"
)

type coachInteractor struct {
	CoachRepository repository.CoachRepository
	CoachPresenter  presenter.CoachPresenter
}

type CoachInteractor interface {
	Get(c []*model.Coach) ([]*model.Coach, error)
	GetById(id string) (*model.Coach, error)
}

// Returns a CoachInteractor
func NewCoachInteractor(r repository.CoachRepository, p presenter.CoachPresenter) CoachInteractor {
	return &coachInteractor{r, p}
}

// Calls FindAll method from CoachRepository and handles its response with ResponseCoaches from CoachPresenter
func (cs *coachInteractor) Get(c []*model.Coach) ([]*model.Coach, error) {
	c, err := cs.CoachRepository.FindAll(c)
	if err != nil {
		return nil, err
	}

	return cs.CoachPresenter.ResponseCoaches(c), nil
}

// Calls FindById method from CoachRepository and handles its response with ResponseCoach from CoachPresenter
func (cs *coachInteractor) GetById(id string) (*model.Coach, error) {
	c, err := cs.CoachRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return cs.CoachPresenter.ResponseCoach(c), nil
}
