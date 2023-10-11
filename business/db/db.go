package db

import (
	amenitieClient "mvc-go/clients/amenitie"
	bookingClient "mvc-go/clients/booking"
	businessClient "mvc-go/clients/business"
	hotelClient "mvc-go/clients/hotel"
	photoClient "mvc-go/clients/photo"
	userClient "mvc-go/clients/user"
	"os"

	"mvc-go/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {

	// DB Connections Parameters
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

	// We need to add all CLients that we built
	userClient.Db = db
	hotelClient.Db = db
	photoClient.Db = db
	amenitieClient.Db = db
	bookingClient.Db = db
	businessClient.Db = db

}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate((&model.User{}))
	db.AutoMigrate((&model.Hotel{}))
	db.AutoMigrate((&model.Booking{}))
	db.AutoMigrate((&model.Photo{}))
	db.AutoMigrate((&model.Amenitie{}))
	db.AutoMigrate((&model.HotelMapping{}))

	log.Info("Finishing Migration Database Tables")
}
