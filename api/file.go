package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jgero/schlingel/file"
	"github.com/jgero/schlingel/model"
)

type Persistence interface {
	CreateFile(file model.SchlingelFile) error
	//DeleteFile(file model.SchlingelFile) error
	//DeleteFileByUuid(uuid.UUID) error
	//UpdateFile(file model.SchlingelFile) error
	GetFileByUuid(uuid.UUID) (*model.SchlingelFile, error)
}

func routerAddUploadFile(router *gin.Engine, persistence Persistence) {
	router.POST("/files/upload", func(c *gin.Context) {
		uploadedFile, err := c.FormFile("file")
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		filetype, err := file.GetFileTypeFromFilename(uploadedFile.Filename)
		if err != nil {
			c.String(http.StatusBadRequest, "Unsupported filetype")
			return
		}

		fileUuid := uuid.New()
		file := model.SchlingelFile{UUID: fileUuid, Filename: uploadedFile.Filename, FileHash: "", FileType: filetype, OwnerUUID: uuid.New()}

		persistence.CreateFile(file)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.SaveUploadedFile(uploadedFile, fmt.Sprintf("./%s", uploadedFile.Filename))

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded! UUID: %s", uploadedFile.Filename, file.UUID.String()))
	})
}

func routerAddDownloadFile(router *gin.Engine, persistence Persistence) {
	router.GET("/files/:uuid", func(c *gin.Context) {
		fileUuidStr := c.Param("uuid")

		fileUuid, err := uuid.Parse(fileUuidStr)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid uuid")
		}

		file, err := persistence.GetFileByUuid(fileUuid)
		if file == nil {
			c.Status(http.StatusNotFound)
			return
		}

		filepath := fmt.Sprintf("./%s", file.Filename)

		if _, err := os.Stat(filepath); err != nil {
			// this shouldn't happen
			c.String(http.StatusNotFound, fmt.Sprintf("'%s' does not exist.", file.Filename))
			return
		}

		c.File(filepath)

		c.String(http.StatusOK, fmt.Sprintf("'%s' downloaded!", file.Filename))
	})
}
