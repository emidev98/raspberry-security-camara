package main

import (
	"github.com/emidev98/raspberry-security-camara/services"
)

const videoOutputFolder = "records"

func main() {
	router := services.NewRouterService(videoOutputFolder)
	camService := services.NewCamaraService(videoOutputFolder)

	go func() {
		camService.StartRecording()
	}()

	router.InitRestRouter()
}
