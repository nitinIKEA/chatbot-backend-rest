package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Service) FaqDetailedfaqGet(w http.ResponseWriter, r *http.Request) {
	var message string
	vars := mux.Vars(r)
	switch vars["detailedfaq"] {
	case "Create Showroom Order":
		message = `How to create a showroom order?
		For more details refer below:
		https://confluence.build.ingka.ikea.com/display/ILO/.DD+UC919+Create+Showroom+Order+vBase`
	case "Return order":
		message = `How to create a return order?
		For more details refer below:
		https://confluence.build.ingka.ikea.com/display/ILO/.DD+UC903+Create+Return+Order+vBase`
	case "Clearing order":
		message = `How to create a clearing order?
		For more details refer below:
		https://confluence.build.ingka.ikea.com/display/ILO/.DD+UC902+Create+Clearing+Order+vBase`
	default:
		message = "Invalid selection"
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
