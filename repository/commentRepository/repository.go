package commentRepository

import (
	"MyGramHacktiv8/entity"
	"MyGramHacktiv8/pkg/errs"
)

type CommentRepository interface {
	PostComment(comment *entity.Comment) (*entity.Comment, errs.MessageErr)
	GetAllComments() ([]*entity.Comment, errs.MessageErr)
	GetCommentByID(commentID uint) (*entity.Comment, errs.MessageErr)
	EditCommentData(commentID uint, comment *entity.Comment) (*entity.Comment, errs.MessageErr)
	DeleteComment(commentID uint) errs.MessageErr
}
