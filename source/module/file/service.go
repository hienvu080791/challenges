package file

import (
	"demo_challenges/source/module/storage"
	"demo_challenges/source/module/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type Service interface {
	UploadFile(ctx *gin.Context, file *multipart.FileHeader) interface{}
}
type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) UploadFile(ctx *gin.Context, file *multipart.FileHeader) interface{} {
	tempFile, err := os.CreateTemp("/tmp", "")
	if err != nil {
		return utils.RespModel{
			Code:    1,
			Message: "error",
			Data:    err.Error(),
		}
	}
	defer tempFile.Close()

	src, err := file.Open()
	if err != nil {
		return utils.RespModel{
			Code:    1,
			Message: "error",
			Data:    err.Error(),
		}
	}
	defer src.Close()
	fileType := filepath.Ext(file.Filename)
	size := file.Size
	metadata := storage.FileMetadata{
		Filename:    file.Filename,
		ContentType: fileType,
		Size:        size,
		TempPath:    tempFile.Name(),
		UploadTime:  time.Now(),
		RequestInfo: fmt.Sprintf("RemoteAddr: %s, User-Agent: %s", ctx.Request.RemoteAddr, ctx.Request.Header.Get("User-Agent")),
	}
	if err := storage.CreateFileMetadata(metadata); err != nil {
		return utils.RespModel{
			Code:    1,
			Message: "error",
			Data:    err.Error(),
		}
	}
	return utils.RespModel{
		Code:    0,
		Message: "success",
	}
}
