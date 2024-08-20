package response

import "gin-vue-admin/model/example"

type FilePathResponse struct {
	FilePath string `json:"file_path"`
}

type FileResponse struct {
	File example.ExaFile `json:"file"`
}
