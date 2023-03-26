package services

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type downloadService struct {
	downnloadFolder string
}

func NewDownloadService(downnloadFolder string) *downloadService {
	downloadService := &downloadService{
		downnloadFolder: downnloadFolder,
	}
	return downloadService
}

func (s downloadService) FileHandler(w http.ResponseWriter, r *http.Request) {
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
