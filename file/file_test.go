package file

import (
	"testing"

	"github.com/jgero/schlingel/model"
	"github.com/stretchr/testify/assert"
)

func TestGetFiletypeFromFilename(t *testing.T) {
	filetype, err := GetFileTypeFromFilename("foo.pdf")
	assert.NoError(t, err)
	assert.Equal(t, model.FileTypePDF, filetype)

	filetype, err = GetFileTypeFromFilename("foo.PDF")
	assert.NoError(t, err)
	assert.Equal(t, model.FileTypePDF, filetype)

	_, err = GetFileTypeFromFilename("foo.xlsx")
	assert.Error(t, err)
	assert.Equal(t, "unknown file type", err.Error())
}
