package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/conzmr/academy-go-q32021/usecase"
)

var (
	au = usecase.NewCoachUseCase()
)

// Adress - Interface for Coach Controller
type Coach interface {
	ReadCSVCoach(http.ResponseWriter, *http.Request)
}

type c struct{}

// NewCoachController - The constructor for a controller used at routes
func NewCoachController() Coach {
	return &c{}
}

// ReadCSVCoach - Handler to read the all the Coaches from a csv file
func (*c) ReadCSVCoach(w http.ResponseWriter, r *http.Request) {
	ad, err := au.ReadCSVCoach("")

	if err != nil {
		HandleError(w, r, err)
	}

	ja, err := json.Marshal(ad)

	if err != nil {
		HandleError(w, r, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ja)
}

// HandleError - Refactored func to report the errors in the controllers
func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	log.Fatalln(err)
}
