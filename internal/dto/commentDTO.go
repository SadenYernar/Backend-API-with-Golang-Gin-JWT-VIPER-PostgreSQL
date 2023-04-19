package dto

import "Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/model"

type CommentResponceDTO struct {
	PostID       int    `json:"postid"`
	Author       string `json:"author"`
	Content      string `json:"content"`
	LikeField    int    `json:"likeField "`
	DislikeField int    `json:"dislikeField"`
	CreatedAt    string `json:"createdat"`
}

func CommentDTO(comment model.Comment) *CommentResponceDTO {
	return &CommentResponceDTO{
		PostID:       comment.PostID,
		Author:       comment.Author,
		Content:      comment.Content,
		LikeField:    comment.LikeField,
		DislikeField: comment.DislikeField,
		CreatedAt:    comment.CreatedAt,
	}
}
