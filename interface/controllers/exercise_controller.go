package controller

import (
	"net/http"
	"strconv"

	"github.com/conzmr/academy-go-q32021/domain/model"
	"github.com/conzmr/academy-go-q32021/usecase/interactor"
	"github.com/labstack/echo"
)

type exerciseController struct {
	exerciseInteractor interactor.ExerciseInteractor
}

type ExerciseController interface {
	GetExercises(c Context) error
	SyncExercises(c Context) error
}

// Returns ExerciseController
func NewExerciseController(e interactor.ExerciseInteractor) ExerciseController {
	return &exerciseController{e}
}

// Validate query params from given context and calls ExerciseInteractor Get method
func (ec *exerciseController) GetExercises(c Context) error {
	idType := c.QueryParam("type")
	items := c.QueryParam("items")
	itemsPerWorkers := c.QueryParam("items_per_workers")
	var e []*model.Exercise
	i, iErr := strconv.Atoi(items)
	ipw, ipwErr := strconv.Atoi(itemsPerWorkers)
	if idType != "odd" && idType != "even" {
		return echo.NewHTTPError(http.StatusBadRequest, "type only support \"odd\" or \"even\"")
	}
	if i < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "required number of items")
	}
	if ipw < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "required number of items per workers")
	}
	if i < ipw {
		return echo.NewHTTPError(http.StatusBadRequest, "number of items must be greater or equal than number of items per workers")
	}
	if iErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, iErr.Error())
	}
	if ipwErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ipwErr.Error())
	}
	e, err := ec.exerciseInteractor.Get(e, idType, i, ipw)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, e)
}

// Calls ExerciseInteractor Sync method
func (ec *exerciseController) SyncExercises(c Context) error {
	var e []*model.Exercise

	e, err := ec.exerciseInteractor.Sync(e)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, e)
}
