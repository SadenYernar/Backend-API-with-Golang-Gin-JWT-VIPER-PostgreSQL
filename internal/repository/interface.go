package repository

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/model"
	"database/sql"
)

type Repository struct {
	User
	Post
	Comment
	Reaction
	Session
}

type User interface {
	CreateUser(user model.User) (int, model.User, error)
	GetUserInfoByEmail(user model.User) (model.User, error)
	GetUserInfoByUsername(username string) (model.User, error)
	SetSession(user model.User, token string) error
}

type Post interface {
	CreatePost(post model.Post) (int, model.Post, error)
	GetAllPost() (int, []model.Post, error)
	GetUserPosts(username string) (int, []model.Post, error)
}

type Comment interface{}

type Reaction interface{}

type Session interface{}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		User:    NewUserRepository(db),
		Post:    NewPostRepository(db),
		Comment: NewCommentRepository(db),
	}
}
