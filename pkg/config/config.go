package config

import "time"

const (
	PartSize          = 5 * 1024 * 1024
	PreviewTtl        = 24 * time.Hour
	ProfilePictureTtl = time.Hour

	MaxTitleLength       = 100
	MaxDescriptionLength = 5000
	MaxPreviewSize       = 2 * 1024 * 1024
)
