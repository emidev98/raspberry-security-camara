package services

import (
	"os/exec"
	"time"
)

// More infor on how to use the libcamera-vid command can be found here:
// https://www.raspberrypi.com/documentation/computers/camera_software.html
type camaraService struct {
	outputFolder string // Folder where to store the files generated by the camara
	time         string // The default value is "5000" (5 seconds). The value zero causes the application to run indefinitely.

	width     string // The width of the video stream e.g. "1280"
	height    string // The height of the video stream e.g. "720"
	framerate string // The framerate of the video stream e.g. "15"

	preview   string // "1" = the preview window will be shown. "0" = the preview window will be hidden
	autofocus string // "1" = the autofocus will be enabled. "0" = autofocus will be disabled
}

func NewCamaraService(outputFolder string) *camaraService {
	return &camaraService{
		outputFolder: outputFolder,
		width:        "1280",
		height:       "720",
		preview:      "0",
		framerate:    "15",
		autofocus:    "1",
		time:         "0",
	}
}

func (s *camaraService) StartRecording() error {
	for {
		currentTime := time.Now()
		fileName := s.outputFolder + "/" + currentTime.Format(time.RFC3339) + ".mp4"

		cmd := exec.Command(
			"libcamera-vid",
			"--width", s.width,
			"--height", s.height,
			"--qt-preview", s.preview,
			"--framerate", s.framerate,
			"--timeout", s.time,
			"--autofocus-mode", s.autofocus,
			"--output", fileName,
			"--codec", "libav",
		)

		err := cmd.Run()
		if err != nil {
			return err
		}
	}
}
