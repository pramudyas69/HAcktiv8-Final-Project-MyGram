package comment_pg

import (
	"MyGramHacktiv8/entity"
	"MyGramHacktiv8/pkg/errs"
	"MyGramHacktiv8/repository/commentRepository"
	"fmt"
	"gorm.io/gorm"
	//"text/template/parse"
)

type commentPG struct {
	db *gorm.DB
}

func NewCommentPG(db *gorm.DB) commentRepository.CommentRepository {
	return &commentPG{db: db}
}

func (c *commentPG) PostComment(commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr) {
	comment := entity.Comment{}
	comment.UserID = commentPayload.UserID
	photo := entity.Photo{}

	if err := c.db.Model(photo).Where("id = ?", commentPayload.PhotoID).First(&photo).Error; err != nil {
		fmt.Println(commentPayload.PhotoID)
		return nil, errs.NewNotFoundError("Not found")
	}

	if err := c.db.Model(&comment).Create(&commentPayload).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("went wrong")
	}

	if err := c.db.Last(&comment).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &comment, nil
}

func (c *commentPG) GetAllComments() ([]*entity.Comment, errs.MessageErr) {
	comments := []*entity.Comment{}

	err := c.db.Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return comments, nil
}

func (c *commentPG) GetCommentByID(commentID uint) (*entity.Comment, errs.MessageErr) {
	comment := entity.Comment{}

	if err := c.db.Where("id = ?", commentID).First(&comment).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &comment, nil
}

func (c *commentPG) EditCommentData(commentID uint, commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr) {
	comment := entity.Comment{}
	//fmt.Println("Melihat payload", commentPayload)

	err := c.db.Model(&comment).Where("id = ?", commentID).Updates(&commentPayload).Take(&comment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Comment not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &comment, nil
}

func (c *commentPG) DeleteComment(commentID uint) errs.MessageErr {
	comment := entity.Comment{}

	if err := c.db.Where("id = ?", commentID).Delete(&comment).Error; err != nil {
		return errs.NewNotFoundError("Not found")
	}

	return nil
}
