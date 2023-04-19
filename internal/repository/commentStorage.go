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

func (c *CommentStorage) CreateComments(comment model.Comment) (int, model.Comment, error) {
	query, err := c.db.Prepare(`INSERT INTO comments(postID,content,author,likeField,dislikeField,createdat) VALUES ($1,$2,$3,$4,$5,$6)`)
	if err != nil {
		return 0, model.Comment{}, fmt.Errorf("[CommentStorage]:Error with CreateComments method in repository: %w", err)
	}
	_, err = query.Exec(comment.PostID, comment.Content, comment.Author, comment.LikeField, comment.DislikeField, comment.CreatedAt)
	if err != nil {
		return 0, model.Comment{}, fmt.Errorf("Create comment in repository: %w", err)
	}
	return 200, model.Comment{}, nil
}

func (c *CommentStorage) GetAllComments() (int, []model.Comment, error) {
	stmt := `SELECT id, postID,content,author,likeField,dislikeField,createdat FROM comments`
	query, err := c.db.Prepare(stmt)
	if err != nil {
		return 0, nil, err
	}
	row, err := query.Query()
	if err != nil {
		return 0, nil, fmt.Errorf("[CommentStorage]:Error with GetAllComments method in repository: %w", err)
	}
	temp := model.Comment{}
	allComments := []model.Comment{}
	for row.Next() {
		err := row.Scan(&temp.ID, &temp.PostID, &temp.Content, &temp.Author, &temp.LikeField, &temp.DislikeField, &temp.CreatedAt)
		if err != nil {
			return 0, nil, fmt.Errorf("[CommentStorage]:Error with GetAllComments method in repository: %w", err)
		}
		allComments = append(allComments, temp)
	}
	return 200, allComments, nil
}

func (c *CommentStorage) GetCommentsByID(postID int64) (int, []model.Comment, error) {
	row, err := c.db.Query("SELECT id,postID,content,author,likeField,dislikeField,createdat FROM comments WHERE postID=$1", postID)
	if err != nil {
		return 0, nil, fmt.Errorf("[CommentStorage]:Error with GetCommentsByID method in repository: %w", err)
	}
	temp := model.Comment{}
	allComments := []model.Comment{}
	for row.Next() {
		err := row.Scan(&temp.ID, &temp.PostID, &temp.Content, &temp.Author, &temp.LikeField, &temp.DislikeField, &temp.CreatedAt)
		if err != nil {
			return 0, nil, fmt.Errorf("[CommentStorage]:Error with GetCommentsByID method in repository: %w", err)
		}
		allComments = append(allComments, temp)
	}
	return 200, allComments, nil
}

func (c *CommentStorage) UpdateComment(comment model.Comment) (int, error) {
	stmt := `UPDATE comments SET id = $1, postID = $2, content = $3, author = $4, likeField = $5, dislikeField = $6, createdat = $7 WHERE id == $1`
	query, err := c.db.Prepare(stmt)
	if err != nil {
		return 0, fmt.Errorf("error executing statement %v:\n%v", stmt, err)
	}
	_, err = query.Exec(&comment.ID, &comment.PostID, &comment.Content, &comment.Author, &comment.LikeField, &comment.DislikeField, &comment.CreatedAt)
	if err != nil {
		return 0, fmt.Errorf("error executing statement %v: %v", stmt, err)
	}
	return 200, nil
}

func (c *CommentStorage) GetCommentByCommentID(commentID int) (int, model.Comment, error) {
	stmt := `SELECT id, postID, content, author, likeField, dislikeField, createdat FROM comments WHERE id == $1`
	query, err := c.db.Prepare(stmt)
	if err != nil {
		return 0, model.Comment{}, fmt.Errorf("error executing statement %v: %v", stmt, err)
	}
	var res model.Comment
	err = query.QueryRow(commentID).Scan(&res.ID, &res.PostID, &res.Content, &res.Author, &res.LikeField, &res.DislikeField, &res.CreatedAt)
	if err != nil {
		return 0, model.Comment{}, fmt.Errorf("error executing statement %v: %v", stmt, err)
	}
	return 200, res, nil
}
