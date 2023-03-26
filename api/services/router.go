package services

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type routerService struct {
	healthcheck     *healthcheckService
	filesService    *filesService
	downloadService *downloadService
}

func NewRouterService(filesFolder string) *routerService {
	createFilesFolder(filesFolder)

	rs := &routerService{
		filesService:    NewFilesService(filesFolder),
		downloadService: NewDownloadService(filesFolder),
		healthcheck:     NewHealthcheckService(),
	}

	return rs
}

func createFilesFolder(filesFolder string) {
	if _, err := os.Stat(filesFolder); os.IsNotExist(err) {
		err := os.Mkdir(filesFolder, 0755)
		if err != nil {
			fmt.Printf("Error creating folder: %v\n", err)
			panic(err)
		}

		fmt.Printf("Folder '%s' created successfully\n", filesFolder)
	}
}

func (s routerService) InitRestRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/files", s.filesService.FilesHandler).Methods("GET")
	router.HandleFunc("/files/{id}", s.filesService.FileHandler).Methods("GET")
	router.HandleFunc("/download/files/{id}", s.downloadService.FileHandler).Methods("GET")
	router.HandleFunc("/healthcheck", s.healthcheck.HealthcheckHandler).Methods("GET")

	log.Printf("Starting web server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
