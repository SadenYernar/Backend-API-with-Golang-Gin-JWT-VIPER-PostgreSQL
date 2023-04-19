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
}

type User interface {
	CreateUser(user model.User) (int, model.User, error)
	GetUserInfoByEmail(user model.User) (model.User, error)
	GetUserInfoByUsername(username string) (model.User, error)
}

type Post interface {
	CreatePost(post model.Post) (int, model.Post, error)
	GetPostByID(id int64) (model.Post, error)
	GetAllPost() (int, []model.Post, error)
	GetUserPosts(username string) (int, []model.Post, error)
	GetUsersLikePosts(i []int64) (int, []model.Post, error)
	GetPostsByCategory(category model.Category) (int, []model.Post, error)
	UpdatePost(post model.Post) (int, error)
}

type Comment interface {
	CreateComments(model.Comment) (int, model.Comment, error)
	GetAllComments() (int, []model.Comment, error)
	GetCommentsByID(postID int64) (int, []model.Comment, error)
	GetCommentByCommentID(commentID int) (int, model.Comment, error)
	UpdateComment(comment model.Comment) (int, error)
}

type Reaction interface {
	CreateLikeForPost(like model.LikePost) (model.LikePost, error)
	CreateLikeForComment(like model.LikeComment) (model.LikeComment, error)
	GetUserIDfromLikePost(like model.LikePost) (int64, error)
	GetLikeStatusByPostAndUserID(like model.LikePost) (model.LikeStatus, error)
	GetLikeStatusByCommentAndUserID(like model.LikeComment) (model.LikeStatus, error)
	UpdatePostLikeStatus(like model.LikePost) error
	UpdateCommentLikeStatus(like model.LikeComment) error
	DeletePostLike(model.LikePost) error
	DeleteCommentLike(model.LikeComment) error
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		User:     NewUserRepository(db),
		Post:     NewPostRepository(db),
		Comment:  NewCommentRepository(db),
		Reaction: NewReactionRepository(db),
	}
}
