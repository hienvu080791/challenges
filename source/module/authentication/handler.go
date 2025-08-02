package authentication

import (
	"demo_challenges/source/module/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service Service
}

func New() *Handler {
	return &Handler{
		service: NewService(),
	}
}

func (ins *Handler) Router(route *gin.Engine) {
	route.POST("api/v1/register", ins.register)
	route.POST("api/v1/login", ins.login)
}

func (ins *Handler) register(c *gin.Context) {
	var (
		statusCode = http.StatusOK
		request    RegisterRequest
		response   interface{}
	)
	if err := c.ShouldBindJSON(&request); err != nil {
		response = utils.RespModel{
			Code:    1,
			Message: "Bad Request",
			Data:    err.Error(),
		}
		statusCode = http.StatusBadRequest
	} else {
		response = ins.service.RegisterUser(request.Username, request.Password)
	}
	c.JSON(statusCode, response)
}

func (ins *Handler) login(c *gin.Context) {
	var (
		statusCode = http.StatusOK
		request    LoginRequest
		response   interface{}
	)
	if err := c.ShouldBindJSON(&request); err != nil {
		response = utils.RespModel{
			Code:    1,
			Message: "Bad Request",
			Data:    err.Error(),
		}
		statusCode = http.StatusBadRequest
	} else {
		response = ins.service.LoginUser(request.Username, request.Password)
	}
	c.JSON(statusCode, response)
}
