package core

import "time"

type User struct {
	Id        int       `json:"id" example:"1"`
	Username  string    `json:"username" example:"user123"`
	Email     string    `json:"email" example:"user123@example.com"`
	Password  string    `json:"password" example:"qwerty123"`
	CreatedAt time.Time `json:"created_at" example:"2023-08-25 17:27:35.811169+00"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-08-25 17:27:35.811169+00"`
	DeletedAt time.Time `json:"deleted_at" example:"2023-08-25 17:27:35.811169+00"`
}

func (entity User) IsEmpty() bool {
	return entity == User{}
}
