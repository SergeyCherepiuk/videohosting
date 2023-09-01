package domain

import (
	"time"

	"github.com/google/uuid"
)

type Video struct {
	ID          uuid.UUID `json:"id" db:"id"`
	AuthorID    uuid.UUID `json:"author_id" db:"author_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	PreviewID   uuid.UUID `json:"preview_id" db:"preview_id"`
	ViewsCount  uint      `json:"views_count" db:"views_count"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type VideoService interface {
	Upload(video *Video) error
}
