package services

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type routerService struct {
	healthcheck  *healthcheckService
	filesService *filesService
}

func NewRouterService(videoOutputFolder string) *routerService {
	return &routerService{
		filesService: NewFilesService(videoOutputFolder),
		healthcheck:  NewHealthcheckService(),
	}
}
func (s routerService) InitRestRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", s.healthcheck.HealthcheckHandler).Methods("GET")
	router.HandleFunc("/files", s.filesService.FilesHandler).Methods("GET")
	router.HandleFunc("/files/{id}", s.filesService.FileHandler).Methods("GET")

	log.Printf("Starting web server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
