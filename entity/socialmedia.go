package entity

import "MyGramHacktiv8/dto"

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null;type:varchar(191)" form:"name" json:"name" valid:"required~ame is required"`
	SocialMediaURL string `gorm:"not null" form:"social_media_url" json:"social_media_url" valid:"required~Social Media URL is required"`
	UserID         uint
	User           *User
}

func (s *SocialMedia) ToGetSocialMediaResponseDTO() *dto.GetSocialMediaResponse {
	return &dto.GetSocialMediaResponse{
		ID:             s.ID,
		Name:           s.Name,
		SocialMediaURL: s.SocialMediaURL,
		UserID:         s.UserID,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
		User: dto.EmbeddedUser{
			ID:       s.User.ID,
			Username: s.User.Username,
			Email:    s.User.Email,
		},
	}
}
