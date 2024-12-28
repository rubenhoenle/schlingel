package model

import (
	"github.com/google/uuid"
)

type FileType = string

const (
	FileTypePDF  FileType = "PDF"
	FileTypeDOCX FileType = "DOCX"
)

type SchlingelFile struct {
	UUID      uuid.UUID
	Filename  string
	OwnerUUID uuid.UUID
	FileHash  string
	FileType  FileType
}
