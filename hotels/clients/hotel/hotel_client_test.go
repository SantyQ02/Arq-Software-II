package hotelClient

import (
	"mvc-go/model"
	"mvc-go/utils/initializers"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	err error
)

func init() {

	initializers.LoadTestEnv("../../utils/initializers/test.env")
	// DB Connections Paramters
	DBName := os.Getenv("MONGO_DB_NAME")
	DBUser := os.Getenv("MONGO_DB_USER")
	DBPass := os.Getenv("MONGO_DB_PASS")
	DBHost := os.Getenv("MONGO_DB_HOST")
	DBPort := os.Getenv("MONGO_DB_PORT")

	clientOpts := options.Client().ApplyURI("mongodb://"+DBUser+":"+DBPass+"@"+DBHost+":"+DBPort+"/?authSource=admin&authMechanism=SCRAM-SHA-256")
	cli,err := mongo.Connect(context.TODO(), clientOpts)
	client=cli
	if err!=nil{
		return err
	}

	Db = client.Database(DBName) 

}

func initTestClient() {
	HotelClient = &hotelClient{}
}

var hotelID uuid.UUID
var hotelModel = model.Hotel{
	AmadeusID:      "0000",
	Title:         	"Test",
	Description:    "Test desciption",
	PricePerDay:    999,
	CityCode:       "City",
	Photos:         nil,
	Amenities: 		nil,
	Active:         true,
}

func TestInsertHotel(t *testing.T) {
	initTestClient()
	hotel := HotelClient.InsertHotel(hotelModel)

	hotelID = hotel.HotelID
	assert.NotEqual(t, uuid.Nil, hotel.HotelID)
	assert.Equal(t, hotelID, hotel.HotelID)
	assert.Equal(t, hotelModel.AmadeusID, hotel.AmadeusID)
	assert.Equal(t, hotelModel.Title, hotel.Title)
	assert.Equal(t, hotelModel.Description, hotel.Description)
	assert.Equal(t, hotelModel.PricePerDay, hotel.PricePerDay)
	assert.Equal(t, hotelModel.CityCode, hotel.CityCode)
	assert.Equal(t, hotelModel.Photos, hotel.Photos)
	assert.Equal(t, hotelModel.Amenities, hotel.Amenities)
	assert.Equal(t, hotelModel.Active, hotel.Active)
}

func TestGetHotelById(t *testing.T) {
	initTestClient()
	hotel := HotelClient.GetHotelById(hotelID.String())

	assert.NotEqual(t, uuid.Nil, hotel.HotelID)
	assert.Equal(t, hotelID, hotel.HotelID)
	assert.Equal(t, hotelModel.AmadeusID, hotel.AmadeusID)
	assert.Equal(t, hotelModel.Title, hotel.Title)
	assert.Equal(t, hotelModel.Description, hotel.Description)
	assert.Equal(t, hotelModel.PricePerDay, hotel.PricePerDay)
	assert.Equal(t, hotelModel.CityCode, hotel.CityCode)
	assert.Equal(t, hotelModel.Photos, hotel.Photos)
	assert.Equal(t, hotelModel.Amenities, hotel.Amenities)
	assert.Equal(t, hotelModel.Active, hotel.Active)
}

func TestUpdateHotel(t *testing.T) {
	initTestClient()
	var hotelModelUpdate = model.Hotel{
		HotelID:           hotelID,
		AmadeusID:        "1111",
		Title:            "New title",
		PricePerDay:      8888,
		Description:      "New desciption",
		CityCode:         "New City",
		Photos:           nil,
		Amenities:        nil,
		Active:         true,
	}
	hotel := HotelClient.UpdateHotel(hotelModelUpdate)

	assert.NotEqual(t, uuid.Nil, hotel.HotelID)
	assert.Equal(t, hotelID, hotel.HotelID)
	assert.Equal(t, hotelModelUpdate.AmadeusID, hotel.AmadeusID)
	assert.Equal(t, hotelModelUpdate.Title, hotel.Title)
	assert.Equal(t, hotelModelUpdate.Description, hotel.Description)
	assert.Equal(t, hotelModelUpdate.PricePerDay, hotel.PricePerDay)
	assert.Equal(t, hotelModelUpdate.CityCode, hotel.CityCode)
	assert.Equal(t, hotelModelUpdate.Photos, hotel.Photos)
	assert.Equal(t, hotelModelUpdate.Amenities, hotel.Amenities)
	assert.Equal(t, hotelModelUpdate.Active, hotel.Active)
}
func TestUpdateHotelError(t *testing.T) {
	initTestClient()
	var hotelModelUpdate = model.User{
		AmadeusID:        "1111",
		Title:            "New title",
		PricePerDay:      8888,
		Description:      "New desciption",
		CityCode:         "New City",
		Photos:           nil,
		Amenities:        nil,
		Active:         true,
	}
	hotel := HotelClient.UpdateHotel(hotelModelUpdate)

	assert.Equal(t, uuid.Nil, hotel.HotelID)
	assert.Equal(t, model.User{}, hotel)
}
func TestDeleteHotel(t *testing.T) {
	err := HotelClient.DeleteHotel(hotelID.String())

	assert.Nil(t, err)
}