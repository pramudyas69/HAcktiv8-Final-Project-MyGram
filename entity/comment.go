package entity

import "MyGramHacktiv8/dto"

type Comment struct {
	GormModel
	UserID  uint
	PhotoID uint
	Message string `gorm:"not null" form:"message" json:"message" valid:"required~Message is required"`
	User    *User
	Photo   *Photo
}

func (c *Comment) ToGetCommentResponseDTO() *dto.GetCommentResponse {
	return &dto.GetCommentResponse{
		ID:        c.ID,
		Message:   c.Message,
		PhotoID:   c.PhotoID,
		UserID:    c.UserID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		User: dto.EmbeddedUserResponse{
			ID:       c.User.ID,
			Username: c.User.Username,
			Email:    c.User.Email,
		},
		Photo: dto.EmbeddedPhotoResponse{
			ID:       c.Photo.ID,
			Title:    c.Photo.Title,
			Caption:  c.Photo.Caption,
			PhotoURL: c.Photo.PhotoURL,
			UserID:   c.Photo.UserID,
		},
	}
}
