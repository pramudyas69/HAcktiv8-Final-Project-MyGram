package dto

import "time"

type PhotoRequest struct {
	Title    string `form:"title" json:"title" valid:"required~Title is required"`
	Caption  string `form:"caption" json:"caption" valid:"required~Caption is required"`
	PhotoURL string `form:"photo_url" json:"photo_url" valid:"required~Photo URL is required"`
}

type PhotoResponse struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type PhotoUpdateResponse struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url"`
	UserID    uint       `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type EmbeddedUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type GetPhotoResponse struct {
	ID        uint                 `json:"id"`
	Title     string               `json:"title"`
	Caption   string               `json:"caption"`
	PhotoURL  string               `json:"photo_url"`
	UserID    uint                 `json:"user_id"`
	CreatedAt *time.Time           `json:"created_at"`
	UpdatedAt *time.Time           `json:"updated_at"`
	User      EmbeddedUserResponse `json:"User"`
}

type DeletePhotoResponse struct {
	Message string `json:"message"`
}
