package model

type Comment struct {
	ID           int
	PostID       int    `json:"postid"`
	Author       string `json:"author"`
	Content      string `json:"content"`
	LikeField    int    `json:"likeField "`
	DislikeField int    `json:"dislikeField"`
	CreatedAt    string `json:"createdat"`
}
