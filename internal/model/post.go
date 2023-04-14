package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type Post struct {
	Uuid       uuid.UUID `json:"uuid"`
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Author     string    `json:"author"`
	CreatedAt  time.Time
	Categories string `json:"categories"`
}
