package services

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type downloadService struct {
	downnloadFolder string
	tokenService    *tokenService
}

func NewDownloadService(downnloadFolder string, tokenService *tokenService) *downloadService {
	downloadService := &downloadService{
		downnloadFolder: downnloadFolder,
		tokenService:    tokenService,
	}
	return downloadService
}

func (s downloadService) FileHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	if s.tokenService.IsValidToken(token) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		return
	}
	vars := mux.Vars(r)
	fileID := vars["id"]

	filePath := fmt.Sprintf("%s/%s", s.downnloadFolder, fileID)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Serve the file
	http.ServeFile(w, r, filePath)
}
