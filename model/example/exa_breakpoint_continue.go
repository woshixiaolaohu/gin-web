package example

import "gin-vue-admin/global"

// ExaFile 文件结构体
type ExaFile struct {
	global.GVA_MODEL
	FileName     string
	FileMD5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// ExaFileChunk 切片结构体
type ExaFileChunk struct {
	global.GVA_MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
