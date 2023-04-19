package repository

import (
	"Backend-API-with-Golang-Gin-JWT-VIPER-PostgreSQL/internal/model"
	"database/sql"
	"fmt"
)

type ReactionStorage struct {
	db *sql.DB
}

func NewReactionRepository(db *sql.DB) *ReactionStorage {
	return &ReactionStorage{
		db: db,
	}
}

func (r *ReactionStorage) CreateLikeForPost(like model.LikePost) (model.LikePost, error) {
	queryForLike, err := r.db.Prepare(`INSERT INTO likePost(userID,postID, status) VALUES ($1,$2,$3)`)
	if err != nil {
		return like, fmt.Errorf("[ReactionStorage]:Error with CreateLikeForPost method in repository: %w", err)
	}
	_, err = queryForLike.Exec(like.UserID, like.PostID, like.Status)
	if err != nil {
		return like, fmt.Errorf("[ReactionStorage]:Error with CreateLikeForPost method in repository: %v", err)
	}
	return like, nil
}

func (r *ReactionStorage) CreateLikeForComment(like model.LikeComment) (model.LikeComment, error) {
	queryForLike, err := r.db.Prepare(`INSERT INTO likeComments(userID, commentsID, status) VALUES ($1,$2,$3)`)
	if err != nil {
		return like, fmt.Errorf("[ReactionStorage]:Error with CreateLikeForComment method in repository: %w", err)
	}
	_, err = queryForLike.Exec(like.UserID, like.CommentsID, like.Status)
	if err != nil {
		return like, fmt.Errorf("[ReactionStorage]:Error with CreateLikeForComment method in repository: %v", err)
	}
	return like, nil
}

func (r *ReactionStorage) UpdatePostLikeStatus(like model.LikePost) error {
	records := ("UPDATE likePost SET status = $1 WHERE postID = $2")
	query, err := r.db.Prepare(records)
	if err != nil {
		return fmt.Errorf("[ReactionStorage]:Error with UpdatePostLikeStatus method in repository: %v", err)
	}
	_, err = query.Exec(like.Status, like.PostID)
	if err != nil {
		return fmt.Errorf("[ReactionStorage]:Error with UpdatePostLikeStatus method in repository: %v", err)
	}
	return nil
}

func (r *ReactionStorage) UpdateCommentLikeStatus(like model.LikeComment) error {
	records := ("UPDATE likeComments SET status = $1 WHERE commentsID = $2")
	query, err := r.db.Prepare(records)
	if err != nil {
		return fmt.Errorf("[ReactionStorage]:Error with UpdateCommentLikeStatus method in repository: %v", err)
	}
	_, err = query.Exec(like.Status, like.CommentsID)
	if err != nil {
		return fmt.Errorf("[ReactionStorage]:Error with UpdateCommentLikeStatus method in repository: %v", err)
	}
	return nil
}

func (r *ReactionStorage) GetUserIDfromLikePost(like model.LikePost) (int64, error) {
	row := r.db.QueryRow("SELECT postID FROM likePost WHERE userID=$1", like.UserID)
	temp := model.LikePost{}
	err := row.Scan(&temp.PostID)
	if err != nil {
		return temp.PostID, fmt.Errorf("[ReactionStorage]:Error with GetUserIDfromLikePost method in repository: %v", err)
	}
	return temp.PostID, nil
}

func (r *ReactionStorage) GetLikeStatusByPostAndUserID(like model.LikePost) (model.LikeStatus, error) {
	stmt := `SELECT status FROM likePost WHERE userID == $1 AND postID == $2`
	query, err := r.db.Prepare(stmt)
	if err != nil {
		return model.NoLike, fmt.Errorf("[ReactionStorage]:Error with GetLikeStatusByPostAndUserID method in repository: %v", err)
	}
	res := query.QueryRow(like.UserID, like.PostID)
	var status model.LikeStatus
	err = res.Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.NoLike, nil
		}
		return model.NoLike, fmt.Errorf("[ReactionStorage]:Error with GetLikeStatusByPostAndUserID method in repository: %v", err)
	}
	return status, nil
}

func (r *ReactionStorage) GetLikeStatusByCommentAndUserID(like model.LikeComment) (model.LikeStatus, error) {
	stmt := `SELECT status FROM likeComments WHERE userID == $1 AND commentsID == $2`
	query, err := r.db.Prepare(stmt)
	if err != nil {
		return model.NoLike, fmt.Errorf("[ReactionStorage]:Error with GetLikeStatusByCommentAndUserID method in repository: %v", err)
	}
	res := query.QueryRow(like.UserID, like.CommentsID)
	var status model.LikeStatus
	err = res.Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.NoLike, nil
		}
		return model.NoLike, fmt.Errorf("[ReactionStorage]:Error with GetLikeStatusByCommentAndUserID method in repository: %v", err)
	}
	return status, nil
}

func (r *ReactionStorage) DeleteCommentLike(like model.LikeComment) error {
	stmt := `DELETE FROM likeComments WHERE commentsID == $1 AND userID == $2`
	query, err := r.db.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = query.Exec(like.CommentsID, like.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReactionStorage) DeletePostLike(like model.LikePost) error {
	stmt := `DELETE FROM likePost WHERE postID == $1 AND userID == $2`
	query, err := r.db.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = query.Exec(like.PostID, like.UserID)
	if err != nil {
		return err
	}
	return nil
}
