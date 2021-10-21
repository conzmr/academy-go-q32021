package registry

import (
	controller "github.com/conzmr/academy-go-q32021/interface/controllers"
	cp "github.com/conzmr/academy-go-q32021/interface/presenters"
	cr "github.com/conzmr/academy-go-q32021/interface/repository"
	ir "github.com/conzmr/academy-go-q32021/interface/repository"
	"github.com/conzmr/academy-go-q32021/usecase/interactor"
)

// Returns a NewCoachController from interface
func (r *registry) NewCoachController() controller.CoachController {
	return controller.NewCoachController(r.NewCoachInteractor())
}

// Returns a NewCoachInteractor with a repository and presenter
func (r *registry) NewCoachInteractor() interactor.CoachInteractor {
	return interactor.NewCoachInteractor(r.NewCoachRepository(), r.NewCoachPresenter())
}

// Returns a NewCoachRepository from interface
func (r *registry) NewCoachRepository() cr.CoachRepository {
	return ir.NewCoachRepository()
}

// Returns a NewCoachPresenter from interface
func (r *registry) NewCoachPresenter() cp.CoachPresenter {
	return cp.NewCoachPresenter()
}
