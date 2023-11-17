package client

import (
	solrDB "mvc-go/solr"
	"mvc-go/dto"
	"mvc-go/utils/initializers"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)


func init() {
	initializers.LoadTestEnv("../.env")
	solrDB.StartTestSolr()
	SolrClient = &solrClient{}

}

var photosDto = dto.Photos{
	dto.Photo{
		PhotoID: uuid.New(),
		Url: "test.url",
	},
}

var amenity1 = dto.Amenity{
	AmenityID: uuid.New(),
	Title:     "Piscina",
}

var amenity2 = dto.Amenity{
	AmenityID: uuid.New(),
	Title:     "Gimnasio",
}

// Crear una lista de comodidades
var amenitiesDto = dto.Amenities{amenity1, amenity2}

var hotel_id_0 = uuid.New()
var hotel_id_1 = uuid.New()
var hotel_id_2 = uuid.New()

var hotelDto0 = dto.Hotel {
	HotelID: hotel_id_0,
	AmadeusID: "Test",
	CityCode: "USH",
	Title: "Hotel 0",
	Description: "Description",
	PricePerDay: 10,
	Photos: nil,
	Amenities: nil,
	Active: true,
}

var hotelDto1 = dto.Hotel {
	HotelID: hotel_id_1,
	AmadeusID: "Test",
	CityCode: "USH",
	Title: "Hotel 1",
	Description: "Description",
	PricePerDay: 10,
	Photos: dto.Photos{},
	Amenities: dto.Amenities{},
	Active: true,
}

var hotelDto2 = dto.Hotel {
	HotelID: hotel_id_2,
	AmadeusID: "Test",
	CityCode: "USH",
	Title: "Hotel 2",
	Description: "Description",
	PricePerDay: 10,
	Photos: photosDto,
	Amenities: amenitiesDto,
	Active: true,
}

var hotelDtoUpdated = dto.Hotel {
	HotelID: hotel_id_1,
	AmadeusID: "Test",
	CityCode: "USH",
	Title: "Hotel 1",
	Description: "edited description",
	PricePerDay: 10,
	Photos: photosDto,
	Amenities: amenitiesDto,
	Active: true,
}

func TestAddHotel(t *testing.T) {
	err0 := SolrClient.AddOrUpdateHotel(hotelDto0)
	err1 := SolrClient.AddOrUpdateHotel(hotelDto1)
	err2 := SolrClient.AddOrUpdateHotel(hotelDto2)

	assert.Equal(t, nil, err0)
	assert.Equal(t, nil, err1)
	assert.Equal(t, nil, err2)
}

func TestUpdateHotel(t *testing.T) {
	err := SolrClient.AddOrUpdateHotel(hotelDtoUpdated)

	assert.Equal(t, nil, err)
}

func TestSearchHotels(t *testing.T){
	hotels, err1 := SolrClient.SearchHotels("USH")
	hotelsFail, err2 := SolrClient.SearchHotels("FFF")

	assert.Equal(t, 3, len(hotels))
	assert.Equal(t, 0, len(hotelsFail))
	assert.Equal(t, nil, err1)
	assert.Equal(t, nil, err2)
}

func TestEmptyCollection(t *testing.T){
	err := SolrClient.EmptyCollection()
	assert.Equal(t, nil, err)
}

func TestGetPhotosFromInterface(t *testing.T) {
    // Caso feliz: La interfaz contiene datos válidos
    input := []interface{}{
        "[{\"photo_id\":\"1a619a01-7685-4693-bab1-ca46b82c4bb7\",\"url\":\"test.url\"}]",
    }
    result := getPhotosFromInterface(input)
    assert.Equal(t, 1, len(result))
	id, _ := uuid.Parse("1a619a01-7685-4693-bab1-ca46b82c4bb7")
    assert.Equal(t, id, result[0].PhotoID)

    // Caso no feliz: La interfaz está vacía
    emptyInput := []interface{}{}
    emptyResult := getPhotosFromInterface(emptyInput)
    assert.Equal(t, 0, len(emptyResult))
}

func TestGetAmenitiesFromInterface(t *testing.T) {
    // Caso feliz: La interfaz contiene datos válidos
    input := []interface{}{
        "[{\"amenity_id\":\"1a619a01-7685-4693-bab1-ca46b82c4bb7\",\"title\":\"Piscina\"}]",
    }
    result := getAmenitiesFromInterface(input)
    assert.Equal(t, 1, len(result))
	id, _ := uuid.Parse("1a619a01-7685-4693-bab1-ca46b82c4bb7")
    assert.Equal(t, id, result[0].AmenityID)

    // Caso no feliz: La interfaz está vacía
    emptyInput := []interface{}{}
    emptyResult := getAmenitiesFromInterface(emptyInput)
    assert.Equal(t, 0, len(emptyResult))
}

func TestGetThumbnailFromInterface(t *testing.T) {
    // Caso feliz: La interfaz contiene una cadena válida
    input := []interface{}{
        "test.thumbnail.url",
    }
    result := getThumbnailFromInterface(input)
    assert.Equal(t, "test.thumbnail.url", result)

    // Caso no feliz: La interfaz está vacía
    emptyInput := []interface{}{}
    emptyResult := getThumbnailFromInterface(emptyInput)
    assert.Equal(t, "", emptyResult)

    // Caso no feliz: La interfaz no contiene una cadena
    invalidInput := []interface{}{
        123, // Un entero en lugar de una cadena
    }
    invalidResult := getThumbnailFromInterface(invalidInput)
    assert.Equal(t, "", invalidResult)
}
