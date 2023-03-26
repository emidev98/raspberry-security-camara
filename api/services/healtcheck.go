package services

import (
	"fmt"
	"net/http"
)

type healthcheckService struct {
}

func NewHealthcheckService() *healthcheckService {
	return &healthcheckService{}
}

func (s healthcheckService) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
