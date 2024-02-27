package request

import "wb-data-service-golang/wb-data-service/internal/module/user/core"

type SignInBody struct {
	Email    string `json:"email" validate:"required,email,max=360"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

func (model SignInBody) ToEntity() core.User {
	return core.User{
		Email:    model.Email,
		Password: model.Password,
	}
}
