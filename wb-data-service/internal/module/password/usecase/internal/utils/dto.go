package utils

import (
	password "wb-data-service-golang/wb-data-service/internal/module/password/core"
	user "wb-data-service-golang/wb-data-service/internal/module/user/core"
)

func ConvertPasswordToUser(entity password.Password) user.User {
	return user.User{
		Id:       entity.UserId,
		Password: entity.Password,
	}
}
