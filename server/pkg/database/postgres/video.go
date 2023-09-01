package postgres

import (
	"github.com/SergeyCherepiuk/videohosting/domain"
	"github.com/SergeyCherepiuk/videohosting/pkg/internal/database"
	"github.com/jmoiron/sqlx"
)

type VideoService struct {
	createStmt *sqlx.NamedStmt
}

func NewVideoService() *VideoService {
	return &VideoService{
		createStmt: database.MustPrepareNamed(db, `insert into videos (id, author_id, title, description, preview_id) values (:id, :author_id, :title, :description, :preview_id)`),
	}
}

func (service VideoService) Upload(video domain.Video) error {
	_, err := service.createStmt.Exec(map[string]any{
		"id":          video.ID,
		"author_id":   video.AuthorID,
		"title":       video.Title,
		"description": video.Description,
		"preview_id":  video.PreviewID,
	})
	return err
}
