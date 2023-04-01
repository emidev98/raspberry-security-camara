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
	tokenService    *tokenService
}

func NewRouterService(filesFolder string) *routerService {
	createFilesFolder(filesFolder)
	tokenService := NewTokenService()

	rs := &routerService{
		filesService:    NewFilesService(filesFolder, tokenService),
		downloadService: NewDownloadService(filesFolder, tokenService),
		healthcheck:     NewHealthcheckService(),
		tokenService:    tokenService,
	}

	rs.tokenService.CreateToeknIfDoesNotExist()

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
	router.HandleFunc("/downloads/{id}", s.downloadService.FileHandler).Methods("GET")
	router.HandleFunc("/tokens", s.tokenService.HandleValidateToken).Methods("POST")
	router.HandleFunc("/healthchecks", s.healthcheck.HealthcheckHandler).Methods("GET")

	log.Printf("Starting web server on :8080")

	log.Fatal(http.ListenAndServeTLS(":8080", "store/cert.pem", "store/key.pem", router))
}
