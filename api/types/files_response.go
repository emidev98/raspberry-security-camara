package types

import (
	"io/fs"
	"time"
)

type FilesResponse = []FileResponse

func NewFilesResponse(files []fs.DirEntry) (fr FilesResponse) {

	for _, file := range files {
		fr = append(fr, NewFileResponse(file.Name(), true))
	}

	return fr
}

func NewFilesResponseWithRanges(files []fs.DirEntry, startDate, endDate time.Time) (fr FilesResponse) {
	for _, file := range files {
		fileResonse := NewFileResponse(file.Name(), true)

		if startDate.IsZero() && endDate.IsZero() {
			fr = append(fr, fileResonse)
		} else if startDate.IsZero() && fileResonse.FileDate.Before(endDate) {
			fr = append(fr, fileResonse)
		} else if endDate.IsZero() && fileResonse.FileDate.After(startDate) {
			fr = append(fr, fileResonse)
		} else if fileResonse.FileDate.After(startDate) && fileResonse.FileDate.Before(endDate) {
			fr = append(fr, fileResonse)
		}
	}

	return fr
}
