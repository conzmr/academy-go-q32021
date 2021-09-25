package main

import (
	"github.com/conzmr/academy-go-q32021/controller"
	"github.com/conzmr/academy-go-q32021/infraestructure"
)

var (
	ac = controller.NewAddressController()
	hr = infraestructure.NewMuxRouter()
)

func main() {
	p := ":3000"
	// Handlers for address
	hr.Get("/coach", ac.ReadCSVAddress)
	hr.Serve(p)
}
