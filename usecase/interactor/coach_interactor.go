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

func NewCoachInteractor(r repository.CoachRepository, p presenter.CoachPresenter) CoachInteractor {
	return &coachInteractor{r, p}
}

func (cs *coachInteractor) Get(c []*model.Coach) ([]*model.Coach, error) {
	c, err := cs.CoachRepository.FindAll(c)
	if err != nil {
		return nil, err
	}

	return cs.CoachPresenter.ResponseCoaches(c), nil
}

func (cs *coachInteractor) GetById(id string) (*model.Coach, error) {
	c, err := cs.CoachRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return cs.CoachPresenter.ResponseCoach(c), nil
}
