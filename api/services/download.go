package services

import (
	"fmt"
	"net/http"
	"os"

	"github.com/emidev98/raspberry-security-camara/types"
	"github.com/gorilla/mux"
)

type downloadService struct {
	downloadFolder string
	tokenService   *tokenService
}

func NewDownloadService(downloadFolder string, tokenService *tokenService) *downloadService {
	downloadService := &downloadService{
		downloadFolder: downloadFolder,
		tokenService:   tokenService,
	}
	return downloadService
}

func (s downloadService) FileHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Token")
	if !s.tokenService.IsValidToken(token) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		return
	}
	vars := mux.Vars(r)
	fileID := vars["id"]

	filePath := fmt.Sprintf("%s/%s", s.downloadFolder, fileID)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Serve the file
	http.ServeFile(w, r, filePath)
}

func (s downloadService) FileLastHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Token")
	if !s.tokenService.IsValidToken(token) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		return
	}
	files, err := os.ReadDir(s.downloadFolder)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading files: %v\n", err)
		return
	}
	latestFile := types.NewFileResponse(
		files[len(files)-1].Name(),
		true,
	)
	filePath := fmt.Sprintf("%s/%s", s.downloadFolder, latestFile.FileId)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Serve the file
	http.ServeFile(w, r, filePath)
}
