package service

import (
	"net/http"
)

func (s *Service) FaqDetailedfaqGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
