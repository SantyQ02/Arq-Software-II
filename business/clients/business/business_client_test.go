package businessClient

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
)

var (
	db           *gorm.DB
	amadeusToken string
	err          error
)

func init() {
	initializers.LoadTestEnv("../../utils/initializers/test.env")
	// DB Connections Paramters
	DBName := os.Getenv("MYSQL_DB_NAME")
	DBUser := os.Getenv("MYSQL_DB_USER")
	DBPass := os.Getenv("MYSQL_DB_PASS")
	DBHost := os.Getenv("MYSQL_DB_HOST")
	DBPort := os.Getenv("MYSQL_DB_PORT")
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":"+DBPort+")/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	Db = db
}

func initTestClient() {
	BusinessClient = &businessClient{}
	go getAmadeusToken()
}

var checkInDate, _ = time.Parse("2006-01-02T15:04:05.000Z", "2022-10-02")
var checkOutDate, _ = time.Parse("2006-01-02T15:04:05.000Z", "2022-10-03")
var hotelID uuid.UUID
var hotelMappingModel = model.HotelMapping{
	AmadeusID: "54DRG5",
}

func TestInsertHotelMapping(t *testing.T) {
	initTestClient()
	err := BusinessClient.InsertHotelMapping(hotelMappingModel)

	assert.Nil(t, err)
}

func TestGetAmadeusIDByHotelID(t *testing.T) {
	initTestClient()
	amadeusID := BusinessClient.GetAmadeusIDByHotelID(hotelID)

	assert.Equal(t, hotelMappingModel.AmadeusID, amadeusID)
}
