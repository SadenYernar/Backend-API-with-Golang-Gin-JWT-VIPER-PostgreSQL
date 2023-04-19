package dto

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/model"
	"time"
)

type PostresponseDTO struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Author     string `json:"author"`
	CreatedAt  time.Time
	Categories model.Category `json:"categories"`
}

func PostDTO(post model.Post) *PostresponseDTO {
	return &PostresponseDTO{
		ID:         post.ID,
		Title:      post.Title,
		Content:    post.Content,
		Author:     post.Author,
		CreatedAt:  post.CreatedAt,
		Categories: post.Categories,
	}
}
