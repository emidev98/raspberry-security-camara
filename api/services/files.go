package services

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type filesService struct {
	outputFolder string
}

func NewFilesService(outputFolder string) *filesService {
	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		err := os.Mkdir(outputFolder, 0755)
		if err != nil {
			fmt.Printf("Error creating folder: %v\n", err)
			panic(err)
		}

		fmt.Printf("Folder '%s' created successfully\n", outputFolder)
	}

	return &filesService{
		outputFolder: outputFolder,
	}
}

func (s filesService) FilesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(s.outputFolder)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading files: %v\n", err)
		return
	}

	for _, file := range files {
		fmt.Fprintln(w, file.Name())
	}
}

func (s filesService) FileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid file ID: %v\n", err)
		return
	}

	files, err := os.ReadDir(s.outputFolder)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading files: %v\n", err)
		return
	}

	if fileID < 0 || fileID >= len(files) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "File not found")
		return
	}

	file := files[fileID]
	fmt.Fprint(w, file.Name())
}
