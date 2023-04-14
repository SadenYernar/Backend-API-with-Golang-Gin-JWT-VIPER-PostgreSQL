package repository

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/model"
	"database/sql"
	"fmt"
)

type CommentStorage struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentStorage {
	return &CommentStorage{
		db: db,
	}
}

func (c *CommentStorage) CreateComment(comment model.Comment) (int, model.Comment, error) {
	query, err := c.db.Prepare(`INSERT INTO comments(postID,content,author,like,dislike,createdat) VALUES ($1,$2,$3,$4,$5,$6)`)
	if err != nil {
		return 0, model.Comment{}, fmt.Errorf("[CommentStorage]:Error with CreateComments method in repository: %w", err)
	}
	_, err = query.Exec(comment.PostID, comment.Content, comment.Author, comment.Like, comment.Dislike, comment.CreatedAt)
	if err != nil {
		return 0, model.Comment{}, fmt.Errorf("Create comment in repository: %w", err)
	}
	return 200, model.Comment{}, nil
}
