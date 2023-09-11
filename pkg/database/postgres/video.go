package postgres

import (
	"github.com/SergeyCherepiuk/videohosting/domain"
	"github.com/SergeyCherepiuk/videohosting/pkg/internal/database"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type VideoService struct {
	getByUUIDStmt *sqlx.NamedStmt
	createStmt    *sqlx.NamedStmt
}

func NewVideoService() *VideoService {
	return &VideoService{
		getByUUIDStmt: database.MustPrepareNamed(db, `select * from videos where id = :id`),
		createStmt:    database.MustPrepareNamed(db, `insert into videos values (:id, :author_id, :title, :description, :preview_id, :views_count, :created_at)`),
	}
}

func (service VideoService) GetByUUID(uuid uuid.UUID) (domain.Video, error) {
	video := domain.Video{}
	err := service.getByUUIDStmt.Get(&video, map[string]any{
		"id": uuid,
	})
	return video, err
}

func (service VideoService) Upload(video domain.Video) error {
	_, err := service.createStmt.Exec(map[string]any{
		"id":          video.ID,
		"author_id":   video.AuthorID,
		"title":       video.Title,
		"description": video.Description,
		"preview_id":  video.PreviewID,
		"views_count": video.ViewsCount,
		"created_at":  video.CreatedAt,
	})
	return err
}
