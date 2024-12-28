package file

import (
	"errors"
	"strings"

	"github.com/jgero/schlingel/model"
)

func GetFileTypeFromFilename(filename string) (model.FileType, error) {
	if strings.HasSuffix(filename, model.FileTypePDF) || strings.HasSuffix(filename, strings.ToLower(model.FileTypePDF)) {
		return model.FileTypePDF, nil
	} else if strings.HasSuffix(filename, model.FileTypeDOCX) || strings.HasSuffix(filename, strings.ToLower(model.FileTypeDOCX)) {
		return model.FileTypeDOCX, nil
	}
	return "", errors.New("unknown file type")
}
