package interactor

import (
	"github.com/conzmr/academy-go-q32021/domain/model"

	"github.com/conzmr/academy-go-q32021/usecase/presenter"
	"github.com/conzmr/academy-go-q32021/usecase/repository"
)

type exerciseInteractor struct {
	ExerciseRepository repository.ExerciseRepository
	ExercisePresenter  presenter.ExercisePresenter
}

type ExerciseInteractor interface {
	Get(c []*model.Exercise) ([]*model.Exercise, error)
}

func NewExerciseInteractor(r repository.ExerciseRepository, p presenter.ExercisePresenter) ExerciseInteractor {
	return &exerciseInteractor{r, p}
}

func (ei *exerciseInteractor) Get(e []*model.Exercise) ([]*model.Exercise, error) {
	e, err := ei.ExerciseRepository.FindAll(e)
	if err != nil {
		return nil, err
	}

	return ei.ExercisePresenter.ResponseExercises(e), nil
}
