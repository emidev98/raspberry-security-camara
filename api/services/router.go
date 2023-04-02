package services

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/rs/cors"
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
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, CORS-enabled GoLang server!")
	})

	r.HandleFunc("/api/v1/files", s.filesService.FilesHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/files/{id}", s.filesService.FileHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/downloads/{id}", s.downloadService.FileHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/auth/token", s.tokenService.HandleValidateToken).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/healthcheck", s.healthcheck.HealthcheckHandler).Methods(http.MethodGet)

	// Configure CORS options
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://localhost:3000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Printf("Starting web server on :8080")

	log.Fatal(http.ListenAndServeTLS(":8080", "store/cert.pem", "store/key.pem", handler))
}
