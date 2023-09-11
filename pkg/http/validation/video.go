package validation

import (
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/SergeyCherepiuk/videohosting/pkg/config"
)

type UploadVideoRequest struct {
	Title       string
	Description string
	Preview     *multipart.FileHeader
	Video       *multipart.FileHeader
}

func (request *UploadVideoRequest) Validate() error {
	request.Title = strings.TrimSpace(request.Title)
	if len(request.Title) == 0 {
		return fmt.Errorf("title is too short")
	} else if len(request.Title) > config.MaxTitleLength {
		return fmt.Errorf("title is too long")
	}

	request.Description = strings.TrimSpace(request.Description)
	if len(request.Description) > config.MaxDescriptionLength {
		return fmt.Errorf("description is too long")
	}

	if !strings.HasPrefix(request.Preview.Header.Get("Content-Type"), "image/") {
		return fmt.Errorf("invalid preview file type")
	}
	if request.Preview.Size > config.MaxPreviewSize {
		return fmt.Errorf("preview file is too big")
	}

	if !strings.HasPrefix(request.Video.Header.Get("Content-Type"), "video/") {
		return fmt.Errorf("invalid video file type")
	}

	return nil
}
