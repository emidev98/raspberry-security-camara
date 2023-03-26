package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/emidev98/raspberry-security-camara/types"
	"github.com/gorilla/mux"
)

type filesService struct {
	outputFolder string
}

func NewFilesService(outputFolder string) *filesService {
	fileService := &filesService{
		outputFolder: outputFolder,
	}

	return fileService
}

func (s filesService) FilesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(s.outputFolder)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading files: %v\n", err)
		return
	}

	filesResponse := types.NewFilesResponse(files)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filesResponse)
}

func (s filesService) FileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileID := vars["id"]

	filePath := fmt.Sprintf("%s/%s", s.outputFolder, fileID)

	// Check if the file exists
	_, err := os.Stat(filePath)
	exists := !os.IsNotExist(err)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types.NewFileResponse(fileID, exists))
}
