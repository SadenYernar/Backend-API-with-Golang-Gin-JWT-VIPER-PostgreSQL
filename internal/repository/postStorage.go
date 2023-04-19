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

func (p *PostStorage) GetPostByID(id int64) (model.Post, error) {
	row := p.db.QueryRow("SELECT id,uuid,title,content,author,createdAt,categories, like, dislike FROM post WHERE id=$1", id)
	temp := model.Post{}
	err := row.Scan(&temp.ID, &temp.Uuid, &temp.Title, &temp.Content, &temp.Author, &temp.CreatedAt, &temp.Categories, &temp.LikeField, &temp.DislikeField)
	if err != nil {
		return temp, fmt.Errorf("[PostStorage]:Error with GetPostByID method in repository: %w", err)
	}
	return temp, nil
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

func (p *PostStorage) GetUsersLikePosts(postIdArray []int64) (int, []model.Post, error) {
	result := []model.Post{}
	for j := 0; j < len(postIdArray); j++ {
		temp := model.Post{}
		row := p.db.QueryRow("SELECT id,uuid,title,content,author,createdAt,categories,likeField,dislikeField FROM post WHERE id=$1", postIdArray[j])
		err := row.Scan(&temp.ID, &temp.Uuid, &temp.Title, &temp.Content, &temp.Author, &temp.CreatedAt, &temp.Categories, &temp.LikeField, &temp.DislikeField)
		if err != nil {
			return 0, nil, fmt.Errorf("[ReactionStorage]:Error with GetUsersLikePosts method in repository: %w", err)
		}
		result = append(result, temp)
	}
	return 200, result, nil
}

func (c *PostStorage) GetPostsByCategory(category model.Category) (int, []model.Post, error) {
	stmt := `SELECT id, uuid, title, content, author, createdat, categories, likeField, dislikeField FROM post WHERE categories&$1 != 0`
	query, err := c.db.Prepare(stmt)
	if err != nil {
		return 0, nil, err
	}
	var res []model.Post
	values, err := query.Query(category)
	if err != nil {
		return 0, nil, err
	}
	for values.Next() {
		var post model.Post
		if err := values.Scan(&post.ID, &post.Uuid, &post.Title, &post.Content, &post.Author, &post.CreatedAt, &post.Categories, &post.LikeField, &post.DislikeField); err != nil {
			return 0, nil, err
		}
		res = append(res, post)
	}
	return 200, res, nil
}

func (p *PostStorage) UpdatePost(post model.Post) (int, error) {
	stmt := `UPDATE post SET id=$1,uuid=$2,title=$3,content=$4,author=$5,createdat=$6,categories=$7,like=$8,dislike=$9 WHERE id == $1`
	query, err := p.db.Prepare(stmt)
	if err != nil {
		return 0, err
	}
	_, err = query.Exec(&post.ID, &post.Uuid, &post.Title, &post.Content, &post.Author, &post.CreatedAt, &post.Categories, &post.LikeField, &post.DislikeField)
	if err != nil {
		return 0, err
	}
	return 200, nil
}
