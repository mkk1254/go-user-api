package models

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	Dob  string `json:"dob" validate:"required"`
}
