package media

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/HoanNghi16/Devall_backend/internal/validator"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Service struct {
	repository *Repository
	cld        *cloudinary.Cloudinary
}

// Khởi tạo service
func NewService(repository *Repository, cld *cloudinary.Cloudinary) *Service {
	return &Service{
		repository: repository,
		cld: cld,
	}
}

func (service *Service) UploadMedia(ctx context.Context,request *MediaRequest, userID uint) (string, error) {
	allowed := validator.AllowedVideo
	maxSize := 3<<30
	if request.Type == "image"{
		allowed = validator.AllowedImage
		maxSize = 10<<20
	}

	file, err := validator.ValidateAndOpenFile(request.File, int64(maxSize), allowed)
	if err != nil{
		return "", err
	}

	result, err := service.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: "Devall_medias",
		ResourceType: request.Type,
	} )

	if err != nil{
		return "", fmt.Errorf("upload thất bại: %w", err)
	}

	name := result.OriginalFilename

	if request.Name != ""{
		name = request.Name
	}

	media := Media{
		Name: name,
		PublicID: result.PublicID,
		URL: result.SecureURL,
		Type: result.ResourceType, 
		UploadedByID: userID,
	}

	err = service.repository.CreateMedia(&media)
	if err != nil{
	 	_, desErr :=service.cld.Upload.Destroy(ctx, uploader.DestroyParams{
			PublicID: media.PublicID,
			ResourceType: media.Type,			
		})
		if desErr != nil{
			slog.Error(
				"cleanup cloudinary failed",
				"public_id", media.PublicID,
				"resource_type", media.Type,
				"error", desErr,
			)
		}
		return "", fmt.Errorf("Upload thất bại! %w",err)
	}
	return media.URL, nil
}