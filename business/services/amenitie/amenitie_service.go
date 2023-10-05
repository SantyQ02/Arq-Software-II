package amenitieService

import (
	"errors"
	amenitieClient "mvc-go/clients/amenitie"
	"mvc-go/dto"
	"mvc-go/model"

	e "mvc-go/utils/errors"

	"github.com/google/uuid"
)

type amenitieService struct{}

type amenitieServiceInterface interface {
	GetAmenitieById(id uuid.UUID) (dto.Amenitie, e.ApiError)
	GetAmenities() (dto.Amenities, e.ApiError)
	DeleteAmenitie(id uuid.UUID) e.ApiError
	InsertAmenitie(amenitieDto dto.Amenitie) (dto.Amenitie, e.ApiError)
	UpdateAmenitie(amenitieDto dto.Amenitie) (dto.Amenitie, e.ApiError)
	LoadAmenities(id uuid.UUID, amenitieDto dto.Amenities) e.ApiError
	UnloadAmenities(id uuid.UUID, amenitieDto dto.Amenities) e.ApiError
}

var (
	AmenitieService amenitieServiceInterface
)

func init() {
	AmenitieService = &amenitieService{}
}

func (s *amenitieService) GetAmenitieById(id uuid.UUID) (dto.Amenitie, e.ApiError) {
	idString := id.String()
	amenitie := amenitieClient.AmenitieClient.GetAmenitieById(idString)
	if amenitie.AmenitieID == uuid.Nil {
		return dto.Amenitie{}, e.NewNotFoundApiError("Amenitie not found")

	}

	amenitieDto := dto.Amenitie{
		AmenitieID: amenitie.AmenitieID,
		Title:      amenitie.Title,
	}
	return amenitieDto, nil
}

func (s *amenitieService) GetAmenities() (dto.Amenities, e.ApiError) {
	amenities := amenitieClient.AmenitieClient.GetAmenities()
	if len(amenities) == 0 {
		return dto.Amenities{}, e.NewInternalServerApiError("Error getting amenities from database", errors.New("error in database"))
	}

	var amenitiesDto []dto.Amenitie

	for _, amenitie := range amenities {
		var amenitieDto dto.Amenitie
		amenitieDto.AmenitieID = amenitie.AmenitieID
		amenitieDto.Title = amenitie.Title

		amenitiesDto = append(amenitiesDto, amenitieDto)
	}

	return amenitiesDto, nil
}

func (s *amenitieService) InsertAmenitie(amenitieDto dto.Amenitie) (dto.Amenitie, e.ApiError) {
	amenitie := model.Amenitie{
		Title: amenitieDto.Title,
	}

	amenitie = amenitieClient.AmenitieClient.InsertAmenitie(amenitie)
	if amenitie.AmenitieID == uuid.Nil {
		return dto.Amenitie{}, e.NewInternalServerApiError("Error trying insert new amenitie", errors.New(""))
	}

	amenitieDto.AmenitieID = amenitie.AmenitieID

	return amenitieDto, nil
}

func (s *amenitieService) DeleteAmenitie(id uuid.UUID) e.ApiError {
	idString := id.String()

	err := amenitieClient.AmenitieClient.DeleteAmenitie(idString)
	if err != nil {
		return e.NewInternalServerApiError("Something went wrong deleting amenitie", nil)
	}

	return nil
}

func (s *amenitieService) UpdateAmenitie(amenitieDto dto.Amenitie) (dto.Amenitie, e.ApiError) {
	idString := amenitieDto.AmenitieID.String()
	amenitie := amenitieClient.AmenitieClient.GetAmenitieById(idString)
	if amenitie.AmenitieID == uuid.Nil {
		return dto.Amenitie{}, e.NewNotFoundApiError("Amenitie not found")

	}
	amenitie.Title = amenitieDto.Title

	amenitie = amenitieClient.AmenitieClient.UpdateAmenitie(amenitie)
	if amenitie.AmenitieID == uuid.Nil {
		return dto.Amenitie{}, e.NewInternalServerApiError("Error updating amenitie", nil)

	}

	amenitieDto = dto.Amenitie{
		AmenitieID: amenitie.AmenitieID,
		Title:      amenitie.Title,
	}
	return amenitieDto, nil
}

func (s *amenitieService) LoadAmenities(id uuid.UUID, amenitiesDto dto.Amenities) e.ApiError {
	idString := id.String()

	var amenities model.Amenities

	for _, amenitie := range amenitiesDto {
		var model model.Amenitie

		model.AmenitieID = amenitie.AmenitieID
		model.Title = amenitie.Title

		amenities = append(amenities, model)
	}

	err := amenitieClient.AmenitieClient.LoadAmenities(idString, amenities)
	if err != nil {
		return e.NewInternalServerApiError("Something went wrong load amenities", nil)
	}

	return nil
}

func (s *amenitieService) UnloadAmenities(id uuid.UUID, amenitiesDto dto.Amenities) e.ApiError {
	idString := id.String()

	var amenities model.Amenities

	for _, amenitie := range amenitiesDto {
		var model model.Amenitie

		model.AmenitieID = amenitie.AmenitieID
		model.Title = amenitie.Title

		amenities = append(amenities, model)
	}

	err := amenitieClient.AmenitieClient.UnloadAmenities(idString, amenities)
	if err != nil {
		return e.NewInternalServerApiError("Something went wrong load amenities", nil)
	}

	return nil
}
