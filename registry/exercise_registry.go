package registry

import (
	controller "github.com/conzmr/academy-go-q32021/interface/controllers"
	cp "github.com/conzmr/academy-go-q32021/interface/presenters"
	cr "github.com/conzmr/academy-go-q32021/interface/repository"
	ir "github.com/conzmr/academy-go-q32021/interface/repository"
	"github.com/conzmr/academy-go-q32021/usecase/interactor"
)

// Returns a NewExerciseController
func (r *registry) NewExerciseController() controller.ExerciseController {
	return controller.NewExerciseController(r.NewExerciseInteractor())
}

// Returns a NewExerciseInteractor with a repository and presenter
func (r *registry) NewExerciseInteractor() interactor.ExerciseInteractor {
	return interactor.NewExerciseInteractor(r.NewExerciseRepository(), r.NewExercisePresenter())
}

// Returns a NewExerciseRepository from interface
func (r *registry) NewExerciseRepository() cr.ExerciseRepository {
	return ir.NewExerciseRepository()
}

// Returns a NewExercisePresenter from interface
func (r *registry) NewExercisePresenter() cp.ExercisePresenter {
	return cp.NewExercisePresenter()
}
