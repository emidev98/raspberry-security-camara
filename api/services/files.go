package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

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
	// recover start date and end date query params
	// parse the params to time.RFC3339
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	parsedStartDate := time.Time{}
	parsedEndDate := time.Time{}

	if startDate != "" {
		parsedStartDate, err = time.Parse(
			time.RFC3339,
			startDate,
		)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error parsing start date : Expected format %s: %v\n", time.RFC3339, err)
			return
		}
	}

	if endDate != "" {
		parsedEndDate, err = time.Parse(
			time.RFC3339,
			endDate,
		)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error parsing end date : Expected format %s: %v\n", time.RFC3339, err)
			return
		}
	}

	var filesResponse types.FilesResponse
	if startDate == "" && endDate == "" {
		filesResponse = types.NewFilesResponse(files)
	} else {
		filesResponse = types.NewFilesResponseWithRanges(files, parsedStartDate, parsedEndDate)
	}

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
