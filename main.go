package main

import (
	"academy-go-q32021/academy-go-q32021/registry"
	"academy-go-q32021/academy-go-q32021/router"
	"fmt"
	"log"

	"github.com/labstack/echo"
)

func main() {

	r := registry.NewRegistry()

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
