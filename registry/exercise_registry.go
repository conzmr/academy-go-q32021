package registry

import (
	controller "github.com/conzmr/academy-go-q32021/interface/controllers"
	cp "github.com/conzmr/academy-go-q32021/interface/presenters"
	cr "github.com/conzmr/academy-go-q32021/interface/repository"
	ir "github.com/conzmr/academy-go-q32021/interface/repository"
	"github.com/conzmr/academy-go-q32021/usecase/interactor"
)

func (r *registry) NewExerciseController() controller.ExerciseController {
	return controller.NewExerciseController(r.NewExerciseInteractor())
}

func (r *registry) NewExerciseInteractor() interactor.ExerciseInteractor {
	return interactor.NewExerciseInteractor(r.NewExerciseRepository(), r.NewExercisePresenter())
}

func (r *registry) NewExerciseRepository() cr.ExerciseRepository {
	return ir.NewExerciseRepository()
}

func (r *registry) NewExercisePresenter() cp.ExercisePresenter {
	return cp.NewExercisePresenter()
}
