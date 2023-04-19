package service

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/model"
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/repository"
)

type Service struct {
	User
	Post
	Comment
	Reaction
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
	GetUserPostsService(username string) (int, []model.Post, error)
}

type Comment interface {
	GetAllCommentsInService() (int, []model.Comment, error)
	GetCommentsByIDinService(postID int64) (int, []model.Comment, error)
	CreateCommentsInService(com model.Comment) (int, model.Comment, error)
	CheckCommentInput(model.Comment) error
}

type Reaction interface {
	LikePostService(like model.LikePost) (int, error)
	LikeCommentService(like model.LikeComment) (int, error)
}

func NewService(repo repository.Repository) Service {
	return Service{
		User:     NewUserService(repo),
		Post:     NewPostService(repo),
		Comment:  NewCommentsService(repo),
		Reaction: NewReactionsService(repo.Reaction, repo.Post, repo.Comment),
	}
}
