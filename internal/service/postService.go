package service

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/model"
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/repository"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type PostService struct {
	repo repository.Repository
}

func NewPostService(repo repository.Repository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (p *PostService) CreatePostService(post model.Post) (int, model.Post, error) {
	err := p.CheckPostInput(post)
	if err != nil {
		return http.StatusBadRequest, post, fmt.Errorf("Create post in service: %w", err)
	}

	post.CreatedAt = time.Now()

	return p.repo.CreatePost(post)
}

func (p *PostService) GetAllPostService() (int, []model.Post, error) {
	return p.repo.GetAllPost()
}

func (p *PostService) GetUserPostsService(username string) (int, []model.Post, error) {
	return p.repo.GetUserPosts(username)
}

func (p *PostService) CheckPostInput(post model.Post) error {
	if len(post.Title) == 0 {
		return errors.New("empty title")
	}
	if title := strings.Trim(post.Title, "\r\n "); len(title) == 0 {
		return errors.New("empty title")
	}
	if content := strings.Trim(post.Content, "\r\n "); len(content) == 0 {
		return errors.New("empty title")
	}
	if len(post.Title) > 50 {
		return errors.New("title too long")
	}
	if len(post.Content) == 0 {
		return errors.New("empty content")
	}
	if len(post.Content) > 1000 {
		return errors.New("content too long")
	}
	return nil
}
