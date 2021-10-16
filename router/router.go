package router

import (
	controller "github.com/conzmr/academy-go-q32021/interface/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/coaches", func(context echo.Context) error { return c.Coach.GetCoaches(context) })
	e.GET("/coaches/:id", func(context echo.Context) error { return c.Coach.GetCoach(context) })
	e.GET("/syncExercises", func(context echo.Context) error { return c.Exercise.SyncExercises(context) })
	e.GET("/exercises", func(context echo.Context) error { return c.Exercise.GetExercises(context) })

	return e
}
