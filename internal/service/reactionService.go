package service

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/model"
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/repository"
)

type ReactionsService struct {
	repo        repository.Reaction
	postRepo    repository.Post
	commentRepo repository.Comment
}

func NewReactionsService(repo repository.Reaction, postRepo repository.Post, commentRepo repository.Comment) *ReactionsService {
	return &ReactionsService{
		repo:        repo,
		postRepo:    postRepo,
		commentRepo: commentRepo,
	}
}

func (r *ReactionsService) LikePostService(like model.LikePost) (int, error) {
	status, err := r.repo.GetLikeStatusByPostAndUserID(like)
	if err != nil {
		return 0, err
	}
	post, err := r.postRepo.GetPostByID(like.PostID)
	if err != nil {
		return 0, err
	}
	if status == model.NoLike {
		_, err := r.repo.CreateLikeForPost(like)
		if err != nil {
			return 0, err
		}
		switch like.Status {
		case model.Like:
			post.LikeField += 1
		case model.DisLike:
			post.DislikeField += 1
		}
		return r.postRepo.UpdatePost(post)
	}
	if status == like.Status {
		switch like.Status {
		case model.Like:
			post.LikeField -= 1
		case model.DisLike:
			post.DislikeField -= 1
		}
		if err := r.repo.DeletePostLike(like); err != nil {
			return 0, err
		}
		return r.postRepo.UpdatePost(post)
	}
	if status != like.Status {
		switch like.Status {
		case model.Like:
			post.LikeField += 1
			post.DislikeField -= 1
		case model.DisLike:
			post.LikeField -= 1
			post.DislikeField += 1
		}
		if _, err := r.postRepo.UpdatePost(post); err != nil {
			return 0, err
		}
		return 0, r.repo.UpdatePostLikeStatus(like)
	}
	return 200, nil
}

func (r *ReactionsService) LikeCommentService(like model.LikeComment) (int, error) {
	status, err := r.repo.GetLikeStatusByCommentAndUserID(like)
	if err != nil {
		return 0, err
	}
	_, comment, err := r.commentRepo.GetCommentByCommentID(like.CommentsID)
	if err != nil {
		return 0, err
	}
	if status == model.NoLike {
		if _, err := r.repo.CreateLikeForComment(like); err != nil {
			return 0, err
		}
		switch like.Status {
		case model.Like:
			comment.LikeField += 1
		case model.DisLike:
			comment.DislikeField += 1
		}
		_, err := r.commentRepo.UpdateComment(comment)
		if err != nil {
			return 0, err
		}
		return 200, nil
	}
	if status == like.Status {
		switch like.Status {
		case model.Like:
			comment.LikeField -= 1
		case model.DisLike:
			comment.DislikeField -= 1
		}
		if _, err := r.commentRepo.UpdateComment(comment); err != nil {
			return 0, err
		}
		if err := r.repo.DeleteCommentLike(like); err != nil {
			return 0, err
		}
	}
	if status != like.Status {
		switch like.Status {
		case model.Like:
			comment.LikeField += 1
			comment.DislikeField -= 1
		case model.DisLike:
			comment.LikeField -= 1
			comment.DislikeField += 1
		}
		if _, err := r.commentRepo.UpdateComment(comment); err != nil {
			return 0, err
		}
		return 200, r.repo.UpdateCommentLikeStatus(like)
	}
	return 200, nil
}
