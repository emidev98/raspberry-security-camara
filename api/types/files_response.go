package types

import "io/fs"

type FilesResponse = []FileResponse

func NewFilesResponse(files []fs.DirEntry) (fr FilesResponse) {

	for _, file := range files {
		fr = append(fr, NewFileResponse(file.Name(), true))
	}

	return fr
}
