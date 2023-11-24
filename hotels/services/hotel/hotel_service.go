package hotelService

import (
	hotelClient "mvc-go/clients/hotel"
	"mvc-go/dto"
	"mvc-go/model"
	"errors"
	e "mvc-go/utils/errors"
	"github.com/google/uuid"
	"mvc-go/queue"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type hotelService struct{}

type hotelServiceInterface interface {
	InsertHotel(hotelDto dto.Hotel) (dto.Hotel, e.ApiError)
	UpdateHotel(hotelDto dto.Hotel) (dto.Hotel, e.ApiError)
	DeleteHotel(id uuid.UUID) e.ApiError
	GetHotels() (dto.Hotels, e.ApiError)
	GetHotelById(id uuid.UUID) (dto.Hotel, e.ApiError)
	UploadPhoto(photoDto dto.Photo, hotelId uuid.UUID) (dto.Photo, e.ApiError)
	SendMessage(id uuid.UUID, action string)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

type Message struct{
    HotelID uuid.UUID  `json:"hotel_id"`
    Action  string `json:"action"`
}

func (s *hotelService) SendMessage(id uuid.UUID, action string){

	q := queue.Queue
	ch := queue.Channel

	message := Message{
        HotelID: id,
        Action:  action,
    }

	messageJSON, err := json.Marshal(message)

	err = ch.Publish(
		"",     // Intercambio (exchange) predeterminado
		q.Name, // Nombre de la cola
		false,  // No mandar confirmación
		false,  // No es mandatorio
		amqp.Publishing{
            ContentType: "application/json", // Establece el tipo de contenido a JSON
            Body:        messageJSON,         // Establece el cuerpo del mensaje como JSON
		},
	)
	if err != nil {
		log.Fatalf("Failed to post a message: %s", err)
		return
	}
}

func (s *hotelService) InsertHotel(hotelDto dto.Hotel) (dto.Hotel, e.ApiError) {

	var Photos model.Photos
	
	var Amenities model.Amenities

	for _, amenity := range hotelDto.Amenities {
			var modelAmenity model.Amenity

			modelAmenity.AmenityID = uuid.New().String()
			modelAmenity.Title = amenity.Title

			Amenities = append(Amenities, modelAmenity)
		}

	hotel := model.Hotel{
		HotelID: 	 uuid.New().String(),
		AmadeusID:	 hotelDto.AmadeusID,
		CityCode: 	 hotelDto.CityCode,
		Title:       hotelDto.Title,
		Description: hotelDto.Description,
		PricePerDay: hotelDto.PricePerDay,
		Photos: 	 Photos,
		Amenities:	 Amenities,
		Active:      true,
	}

	hotel = hotelClient.HotelClient.InsertHotel(hotel)
	if hotel.HotelID == "" {
		return dto.Hotel{}, e.NewInternalServerApiError("Error trying to insert hotel", errors.New(""))
	}

	hotelDto.HotelID,_ = uuid.Parse(hotel.HotelID)

	var newAmenities dto.Amenities

	hotelDto.Amenities = newAmenities
	
	for _, amenity := range hotel.Amenities {
		var dtoAmenity dto.Amenity

		dtoAmenity.AmenityID,_ = uuid.Parse(amenity.AmenityID)
		dtoAmenity.Title = amenity.Title

		hotelDto.Amenities = append(hotelDto.Amenities, dtoAmenity)
	}

	err := hotelClient.HotelClient.HotelMapping(hotel)
	if err != nil {
		return dto.Hotel{}, e.NewInternalServerApiError("Error trying to send request", errors.New(""))
	}

	return hotelDto, nil
}

func (s *hotelService) UpdateHotel(hotelDto dto.Hotel) (dto.Hotel, e.ApiError) {

	var Photos model.Photos

	for _, photo := range hotelDto.Photos {
			var modelPhoto model.Photo

			modelPhoto.PhotoID = photo.PhotoID.String()
			modelPhoto.Url = photo.Url

			Photos = append(Photos, modelPhoto)
		}
	
	var Amenities model.Amenities

	for _, amenity := range hotelDto.Amenities {
			var modelAmenity model.Amenity

			modelAmenity.AmenityID = amenity.AmenityID.String()
			modelAmenity.Title = amenity.Title

			Amenities = append(Amenities, modelAmenity)
		}

	hotel := model.Hotel{
		HotelID:     hotelDto.HotelID.String(),
		AmadeusID:	 hotelDto.AmadeusID,
		CityCode: 	 hotelDto.CityCode,
		Title:       hotelDto.Title,
		Description: hotelDto.Description,
		PricePerDay: hotelDto.PricePerDay,
		Photos: Photos,
		Amenities: Amenities,
		Active:      hotelDto.Active,
	}

	hotel = hotelClient.HotelClient.UpdateHotel(hotel)
	if hotel.HotelID == "" {
		return dto.Hotel{}, e.NewInternalServerApiError("Error trying to update hotel", errors.New(""))
	}
	// hotelDto.HotelID,_ = uuid.Parse(hotel.HotelID)

	return hotelDto, nil
}

func (s *hotelService) DeleteHotel(id uuid.UUID) e.ApiError {
	idString := id.String()

	err := hotelClient.HotelClient.DeleteHotel(idString)
	if err != nil {
		return e.NewInternalServerApiError("Something went wrong deleting hotel", nil)
	}

	return nil

}

func (s *hotelService) GetHotels() (dto.Hotels, e.ApiError) {

	hotels := hotelClient.HotelClient.GetHotels()
	if len(hotels) == 0 {
		return dto.Hotels{}, e.NewNotFoundApiError("Hotels not found")
	}

	var hotelsDto dto.Hotels

	for _, hotel := range hotels {

		idhotel,_ := uuid.Parse(hotel.HotelID)

		hotelDto := dto.Hotel{
			HotelID:     idhotel,
			AmadeusID: 	 hotel.AmadeusID,
			CityCode:    hotel.CityCode,
			Title:       hotel.Title,
			Description: hotel.Description,
			PricePerDay: hotel.PricePerDay,
			Active:      hotel.Active,
		}
		for _, photo := range hotel.Photos {
			var dtoPhoto dto.Photo

			dtoPhoto.PhotoID,_ = uuid.Parse(photo.PhotoID)
			dtoPhoto.Url = photo.Url
			hotelDto.Photos = append(hotelDto.Photos, dtoPhoto)
		}
		for _, amenity := range hotel.Amenities {
			var dtoAmenity dto.Amenity

			dtoAmenity.AmenityID,_ = uuid.Parse(amenity.AmenityID)
			dtoAmenity.Title = amenity.Title

			hotelDto.Amenities = append(hotelDto.Amenities, dtoAmenity)
		}

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}


func (s *hotelService) GetHotelById(id uuid.UUID) (dto.Hotel, e.ApiError) {

	idString := id.String()

	hotel := hotelClient.HotelClient.GetHotelById(idString)
	if hotel.HotelID == "" {
		return dto.Hotel{}, e.NewNotFoundApiError("Hotel not found")
	}

	idhotel,_ := uuid.Parse(hotel.HotelID)

	hotelDto := dto.Hotel{
		HotelID:     idhotel,
		AmadeusID: 	 hotel.AmadeusID,
		CityCode:    hotel.CityCode,
		Title:       hotel.Title,
		Description: hotel.Description,
		PricePerDay: hotel.PricePerDay,
		Active:      hotel.Active,
	}
	for _, photo := range hotel.Photos {
		var dtoPhoto dto.Photo

		dtoPhoto.PhotoID,_ = uuid.Parse(photo.PhotoID)
		dtoPhoto.Url = photo.Url
		hotelDto.Photos = append(hotelDto.Photos, dtoPhoto)
	}
	for _, amenity := range hotel.Amenities {
		var dtoAmenity dto.Amenity

		dtoAmenity.AmenityID,_ = uuid.Parse(amenity.AmenityID)
		dtoAmenity.Title = amenity.Title

		hotelDto.Amenities = append(hotelDto.Amenities, dtoAmenity)
	}

	return hotelDto, nil
}

func (s *hotelService) UploadPhoto(photoDto dto.Photo, hotelId uuid.UUID) (dto.Photo, e.ApiError) {
	photo := model.Photo{
		Url:     photoDto.Url,
		PhotoID: uuid.New().String(),
	}

	hotel := hotelClient.HotelClient.GetHotelById(hotelId.String())
	if hotel.HotelID == "" {
		return dto.Photo{}, e.NewInternalServerApiError("Hotel not found", errors.New(""))
	}

	if hotel.Photos != nil {
		hotel.Photos = append(hotel.Photos, photo)
	}else {
		var Photos model.Photos
		Photos = append(Photos, photo)		
		hotel.Photos = Photos
	}

	hotel = hotelClient.HotelClient.UpdateHotel(hotel)
	if hotel.HotelID == "" {
		return dto.Photo{}, e.NewInternalServerApiError("Error trying to insert photo", errors.New(""))
	}

	photoDto.PhotoID,_ = uuid.Parse(photo.PhotoID)

	return photoDto, nil
}

