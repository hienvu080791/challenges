package authentication

import (
	"demo_challenges/source/module/storage"
	"demo_challenges/source/module/utils"
)

type Service interface {
	RegisterUser(username, password string) interface{}
	LoginUser(username, password string) interface{}
}
type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) RegisterUser(username, password string) interface{} {
	err := storage.CreateUser(storage.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		return utils.RespModel{
			Code:    1,
			Message: "error",
			Data:    err.Error(),
		}
	}
	return utils.RespModel{
		Code:    0,
		Message: "success",
		Data:    nil,
	}
}

func (s *service) LoginUser(username, password string) interface{} {
	err := storage.CheckLogin(storage.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		return utils.RespModel{
			Code:    1,
			Message: "error",
			Data:    err.Error(),
		}
	}
	token, err := utils.GenerateToken(username)
	if err != nil {
		return utils.RespModel{
			Code:    1,
			Message: "error",
			Data:    err.Error(),
		}
	}
	return utils.RespModel{
		Code:    0,
		Message: "login success",
		Data: struct {
			Token string `json:"token"`
		}{
			Token: token,
		},
	}
}
