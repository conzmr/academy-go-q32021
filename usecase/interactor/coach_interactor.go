package interactor

import (
	"academy-go-q32021/academy-go-q32021/domain/model"
	"errors"

	"github.com/conzmr/academy-go-q32021/domain/model"
	"github.com/conzmr/academy-go-q32021/usecase/presenter"
	"github.com/conzmr/academy-go-q32021/usecase/repository"
)

type coachInteractor struct {
	CoachRepository repository.CoachRepository
	CoachPresenter  presenter.CoachPresenter
	DBRepository    repository.DBRepository
}

type CoachInteractor interface {
	Get(u []*model.Coach) ([]*model.Coach, error)
	Create(u *model.Coach) (*model.Coach, error)
}

func NewCoachInteractor(r repository.CoachRepository, p presenter.CoachPresenter, d repository.DBRepository) CoachInteractor {
	return &coachInteractor{r, p, d}
}

func (us *coachInteractor) Get(u []*model.Coach) ([]*model.Coach, error) {
	u, err := us.CoachRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.CoachPresenter.ResponseCoaches(u), nil
}

func (c *coachInteractor) Create(u *model.Coach) (*model.Coach, error) {
	data, err := c.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		c, err := c.CoachRepository.Create(c)

		// do mailing
		// do logging
		// do another process
		return c, err
	})
	coach, ok := data.(*model.Coach)

	if !ok {
		return nil, errors.New("cast error")
	}

	if !errors.Is(err, nil) {
		return nil, err
	}

	return coach, nil
}
