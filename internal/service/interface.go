package service

import (
	"project/internal/model"
	"project/internal/repository"
)

type Service struct {
	User
	Post
	Comment
	Reaction
	Session
}

type User interface {
	CreateUserService(user model.User) (int, model.User, error)
	AuthorizationUserService(user model.User) (string, error)
	GetUserInfoServiceByUsername(username string) (model.User, error)
}

type Post interface {
	CreatePostService(post model.Post) (int, model.Post, error)
	CheckPostInput(post model.Post) error
	GetAllPostService() (int, []model.Post, error)
}

type Comment interface{}

type Reaction interface{}

type Session interface {
	DeleteSessionService(user model.User) error
}

func NewService(repo repository.Repository) Service {
	return Service{
		User: NewUserService(repo),
		Post: NewPostService(repo),
	}
}
