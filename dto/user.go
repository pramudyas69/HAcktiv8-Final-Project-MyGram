package dto

import "time"

type LoginRequest struct {
	Email    string `form:"email" json:"email" valid:"required~Email is required,email~Email is not valid"`
	Password string `form:"password" json:"password" valid:"required~Password is required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterRequest struct {
	Username string `json:"username" form:"username" valid:"required~Username is required"`
	Email    string `json:"email" form:"email" valid:"required~Email is required,email~Email is not valid"`
	Password string `json:"password" form:"password" valid:"required~Password is required"`
	Age      uint8  `json:"age" form:"age" valid:"required~Age is required,range(8|100)~Age must be between 8 and 100"`
}

type RegisterResponse struct {
	Age      uint8  `json:"age"`
	Email    string `json:"email"`
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type UpdateUserDataRequest struct {
	Email    string `json:"email" form:"email" valid:"required~Email is required,email~Email is not valid"`
	Username string `json:"username" form:"username" valid:"required~Username is required"`
}

type UpdateUserDataResponse struct {
	ID        uint       `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Age       uint8      `json:"age"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
