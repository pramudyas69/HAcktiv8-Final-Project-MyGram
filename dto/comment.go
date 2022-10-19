package dto

import "time"

type CommentRequest struct {
	Message string `form:"message" json:"message" valid:"required~Message is required"`
	PhotoID uint   `form:"photo_id" json:"photo_id" valid:"required~Photo ID is required"`
}

type CommentResponse struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	PhotoID   uint       `json:"photo_id"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type EmbeddedPhotoResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

type GetCommentResponse struct {
	ID        uint                  `json:"id"`
	Message   string                `json:"message"`
	PhotoID   uint                  `json:"photo_id"`
	UserID    uint                  `json:"user_id"`
	CreatedAt *time.Time            `json:"created_at"`
	UpdatedAt *time.Time            `json:"updated_at"`
	User      EmbeddedUserResponse  `json:"User"`
	Photo     EmbeddedPhotoResponse `json:"Photo"`
}

type EditCommentRequest struct {
	Message string `form:"message" json:"message" valid:"required~Message is required"`
}

type EditCommentResponse struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	PhotoID   uint       `json:"photo_id"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type DeleteCommentResponse struct {
	Message string `json:"message"`
}
