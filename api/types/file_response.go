package types

import "time"

type FileResponse struct {
	FileId   string    `json:"fileId"`
	Exists   bool      `json:"exists"`
	FileDate time.Time `json:"fileDate"`
}

func NewFileResponse(fileId string, exists bool) FileResponse {
	// transfer the file name format 'records/2023-03-26T20:23:06+03:00.mp4'
	// to time.Time format
	fileDate, _ := time.Parse(time.RFC3339, fileId[:len(fileId)-4])

	return FileResponse{
		FileId:   fileId,
		Exists:   exists,
		FileDate: fileDate,
	}
}
