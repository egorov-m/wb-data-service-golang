package request

import "wb-data-service-golang/wb-data-service/internal/module/user/core"

type SignUpBody struct {
	Username string `json:"username" validate:"omitempty,max=255"`
	Email    string `json:"email" validate:"required,email,max=360"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

func (model SignUpBody) ToEntity() core.User {
	return core.User{
		Username: model.Username,
		Email:    model.Email,
		Password: model.Password,
	}
}
