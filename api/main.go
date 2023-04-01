package main

import (
	"github.com/emidev98/raspberry-security-camara/services"
	"github.com/emidev98/raspberry-security-camara/util"
)

const videoOutputFolder = "records"

func main() {
	util.NewCertificateIfDoesNotExist()

	router := services.NewRouterService(videoOutputFolder)
	router.InitRestRouter()

	camService := services.NewCamaraService(videoOutputFolder)
	go func() {
		camService.StartRecording()
	}()
}
