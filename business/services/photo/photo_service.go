package photoService

import (
	"errors"
	photoClient "mvc-go/clients/photo"
	"mvc-go/dto"
	"mvc-go/model"

	// "time"

	e "mvc-go/utils/errors"

	"github.com/google/uuid"
)

type photoService struct{}

type photoServiceInterface interface {
	GetPhotosByHotelId(id uuid.UUID) (dto.Photos, e.ApiError)
	InsertPhoto(photoDto dto.Photo) (dto.Photo, e.ApiError)
	DeletePhoto(id uuid.UUID) e.ApiError
}

var (
	PhotoService photoServiceInterface
)

func init() {
	PhotoService = &photoService{}
}

func (s *photoService) GetPhotosByHotelId(id uuid.UUID) (dto.Photos, e.ApiError) {
	idString := id.String()
	photos := photoClient.PhotoClient.GetPhotosByHotelId(idString)
	if len(photos) == 0 {
		return dto.Photos{}, e.NewNotFoundApiError("Photos not found")
	}

	var photosDto dto.Photos

	for _, photo := range photos {
		var photoDto dto.Photo
		photoDto.PhotoID = photo.PhotoID
		photoDto.Url = photo.Url
		photoDto.HotelID = photo.HotelID

		photosDto = append(photosDto, photoDto)
	}

	return photosDto, nil
}

func (s *photoService) InsertPhoto(photoDto dto.Photo) (dto.Photo, e.ApiError) {
	photo := model.Photo{
		Url:     photoDto.Url,
		HotelID: photoDto.HotelID,
	}

	photo = photoClient.PhotoClient.InsertPhoto(photo)
	if photo.PhotoID == uuid.Nil {
		return dto.Photo{}, e.NewInternalServerApiError("Error trying insert new photo", errors.New(""))
	}

	photoDto.PhotoID = photo.PhotoID

	return photoDto, nil
}

func (s *photoService) DeletePhoto(id uuid.UUID) e.ApiError {
	idString := id.String()

	err := photoClient.PhotoClient.DeletePhoto(idString)
	if err != nil {
		return e.NewInternalServerApiError("Something went wrong deleting photo", nil)
	}

	return nil
}
