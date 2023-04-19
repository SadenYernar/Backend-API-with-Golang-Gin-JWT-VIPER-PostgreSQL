package service

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/model"
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/repository"
	"errors"
	"strings"
	"time"
)

type CommentService struct {
	repo repository.Comment
}

func NewCommentsService(repo repository.Comment) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (c *CommentService) CheckCommentInput(comment model.Comment) error {
	if comment := strings.Trim(comment.Content, "\r\n "); len(comment) == 0 {
		return errors.New("empty title")
	}
	if len(comment.Content) == 0 {
		return errors.New("empty comment")
	}
	if len(comment.Content) > 500 {
		return errors.New("comment too long")
	}
	return nil
}

func (c *CommentService) GetAllCommentsInService() (int, []model.Comment, error) {
	return c.repo.GetAllComments()
}

func (c *CommentService) GetCommentsByIDinService(postID int64) (int, []model.Comment, error) {
	return c.repo.GetCommentsByID(postID)
}

func (c *CommentService) CreateCommentsInService(com model.Comment) (int, model.Comment, error) {
	t := time.Now()
	timeFormat := t.Format("15:04:04,02 January 2006")
	com.CreatedAt = timeFormat
	return c.repo.CreateComments(com)
}
