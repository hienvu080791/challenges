package file

import (
	"demo_challenges/source/module/middleware"
	"demo_challenges/source/module/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service    Service
	middleware middleware.Service
}

func New() *Handler {
	return &Handler{
		service:    NewService(),
		middleware: middleware.NewService(),
	}
}

func (ins *Handler) Router(route *gin.Engine) {
	route.POST("api/v1/upload", ins.middleware.Middleware(true, ins.uploadFile))

}

func (ins *Handler) uploadFile(c *gin.Context) {
	var (
		statusCode    = http.StatusOK
		response      interface{}
		maxUploadSize = 8 * 1024 * 1024 // 8 MB
	)

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(maxUploadSize))
	if err := c.Request.ParseMultipartForm(int64(maxUploadSize)); err != nil {
		statusCode = http.StatusRequestEntityTooLarge
		response = utils.RespModel{
			Code:    1,
			Message: "error",
			Data:    "File size exceeds the limit of 8MB",
		}
		c.JSON(statusCode, response)
		return
	}

	if len(c.Request.MultipartForm.File) == 0 {
		statusCode = http.StatusBadRequest
		response = utils.RespModel{
			Code:    1,
			Message: "error",
			Data:    "No file uploaded",
		}
		c.JSON(statusCode, response)
		return
	}

	file, err := c.FormFile("data")
	if err != nil {
		statusCode = http.StatusBadRequest
		response = utils.RespModel{
			Code:    1,
			Message: "error",
			Data:    err.Error(),
		}
		c.JSON(statusCode, response)
		return
	}
	response = ins.service.UploadFile(c, file)
	c.JSON(200, response)
}
