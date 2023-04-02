package services

import (
	"encoding/json"
	"net/http"

	"github.com/emidev98/raspberry-security-camara/types"
)

type healthcheckService struct {
}

func NewHealthcheckService() *healthcheckService {
	return &healthcheckService{}
}

func (s healthcheckService) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types.NewHealthcheck())
}
