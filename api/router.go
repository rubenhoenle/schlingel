package api

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/rubenhoenle/schlingel/page"
)

func BuildRouter(persistence Persistence) *gin.Engine {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	routerAddUploadFile(router, persistence)
	routerAddDownloadFile(router, persistence)
	router.GET("/", func(c *gin.Context) {
		_ = render(c, http.StatusOK, page.Index())
	})
	return router
}

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}
