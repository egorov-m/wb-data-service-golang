package model

import (
	"database/sql"
	"wb-data-service-golang/wb-data-service/internal/module/user/core"
)

type User struct {
	Id        int            `db:"id"`
	Username  sql.NullString `db:"username"`
	Email     string         `db:"email"`
	Password  string         `db:"password_hash"`
	CreatedAt sql.NullTime   `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
	DeletedAt sql.NullTime   `db:"deleted_at"`
}

func NewUser(entity core.User) User {
	return User{
		Id: entity.Id,
		Username: sql.NullString{
			String: entity.Username,
			Valid:  entity.Username != "",
		},
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedAt: sql.NullTime{Time: entity.CreatedAt, Valid: true},
		UpdatedAt: sql.NullTime{Time: entity.UpdatedAt, Valid: true},
		DeletedAt: sql.NullTime{Time: entity.DeletedAt, Valid: true},
	}
}

func (model User) TableName() string {
	return "user"
}

func (model User) ToEntity() core.User {
	return core.User{
		Id:        model.Id,
		Username:  model.Username.String,
		Email:     model.Email,
		Password:  model.Password,
		CreatedAt: model.CreatedAt.Time,
		UpdatedAt: model.UpdatedAt.Time,
		DeletedAt: model.DeletedAt.Time,
	}
}
