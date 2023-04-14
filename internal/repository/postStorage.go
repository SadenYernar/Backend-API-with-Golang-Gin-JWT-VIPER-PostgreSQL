package repository

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/model"
	"database/sql"
	"fmt"
	"net/http"
)

type PostStorage struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostStorage {
	return &PostStorage{
		db: db,
	}
}

func (p *PostStorage) CreatePost(post model.Post) (int, model.Post, error) {
	query, err := p.db.Prepare(`INSERT INTO post (uuid,title,content,author,createdat,categories) VALUES ($1,$2,$3,$4,$5,$6)`)
	if err != nil {
		return 0, post, fmt.Errorf("[PostStorage#1]:Error with CreatePost method in repository: %w", err)
	}

	_, err = query.Exec(post.Uuid, post.Title, post.Content, post.Author, post.CreatedAt, post.Categories)
	if err != nil {
		return 0, post, fmt.Errorf("[PostStorage#2]:Error with CreatePost method in repository: %w", err)
	}

	fmt.Println("Post created successfully!")
	return http.StatusOK, post, nil
}

func (p *PostStorage) GetAllPost() (int, []model.Post, error) {
	row, err := p.db.Query("SELECT id,uuid,title,content,author,createdAt,categories FROM post")
	if err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("[PostStorage]:Error with GetAllPost method in repository: %w", err)
	}
	temp := model.Post{}
	allPost := []model.Post{}
	for row.Next() {
		err := row.Scan(&temp.ID, &temp.Uuid, &temp.Title, &temp.Content, &temp.Author, &temp.CreatedAt, &temp.Categories)
		if err != nil {
			return http.StatusInternalServerError, nil, fmt.Errorf("[PostStorage]:Error with GetAllPost method in repository: %w", err)
		}
		allPost = append(allPost, temp)
	}
	return http.StatusOK, allPost, nil
}

func (p *PostStorage) GetUserPosts(username string) (int, []model.Post, error) {
	row, err := p.db.Query("SELECT id,uuid,title,content,author,createdAt,categories FROM post WHERE author=$1", username)
	if err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("[PostStorage]:Error with GetAllPost method in repository: %w", err)
	}
	temp := model.Post{}
	allPost := []model.Post{}
	for row.Next() {
		err := row.Scan(&temp.ID, &temp.Uuid, &temp.Title, &temp.Content, &temp.Author, &temp.CreatedAt, &temp.Categories)
		if err != nil {
			return http.StatusInternalServerError, nil, fmt.Errorf("[PostStorage]:Error with GetAllPost method in repository: %w", err)
		}
		allPost = append(allPost, temp)
	}
	return http.StatusOK, allPost, nil
}
