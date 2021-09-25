package registry

import (
	controller "academy-go-q32021/academy-go-q32021/interface/controllers"
	cp "academy-go-q32021/academy-go-q32021/interface/presenter"
	cr "academy-go-q32021/academy-go-q32021/interface/repository"
	ir "academy-go-q32021/academy-go-q32021/interface/repository"
	"academy-go-q32021/academy-go-q32021/usecase/interactor"
)

func (r *registry) NewUserController() controller.CoachController {
	return controller.NewCoachController(r.NewCoachInteractor())
}

func (r *registry) NewCoachInteractor() interactor.CoachInteractor {
	return interactor.NewCoachInteractor(r.NewCoachRepository(), r.NewCoachPresenter())
}

func (r *registry) NewCoachRepository() cr.CoachRepository {
	return ir.NewCoachRepository(r.db)
}

func (r *registry) NewCoachPresenter() cp.CoachPresenter {
	return cp.NewCoachPresenter()
}
