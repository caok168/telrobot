package models

type CreateUserRequest struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type UpdateUserRequest struct {
	Id                   int    `json:"id"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type ListUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Page  int    `json:"page"`
	Count int    `json:"count"`
}
