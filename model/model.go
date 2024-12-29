package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileType = string

const (
	FileTypePDF  FileType = "PDF"
	FileTypeDOCX FileType = "DOCX"
)

type SchlingelFile struct {
	gorm.Model
	UUID      uuid.UUID
	Filename  string
	OwnerUUID uuid.UUID
	FileHash  string
	FileType  FileType
}
