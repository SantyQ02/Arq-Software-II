package hotelController

import (
	// "mvc-go/model"
	// userService "mvc-go/services/user"
	"net/http"
	"path/filepath"
	"encoding/json"

	"net/http"
	"mvc-go/dto"
	hotelService "mvc-go/services/hotel"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"mvc-go/queue"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct{
    HotelID uuid.UUID  `json:"hotel_id"`
    Action  string `json:"action"`
}

type Bodystruct struct {
    HotelID uuid.UUID `json:"hotel_id,omitempty"`
	AmadeusID  string `json:"amadeus_id,omitempty"`
}

func SendMessage(id uuid.UUID, action string){

	q := queue.Queue
	ch := queue.Channel

	message := Message{
        HotelID: id,
        Action:  action,
    }

	messageJSON, err := json.Marshal(message)

    if err != nil {
        log.Fatalf("Failed to marshal json: %v", err)
        return
    }

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


func GetHotelById(c *gin.Context) {
	uuid, err := uuid.Parse(c.Param("hotelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "HotelID must be a uuid"})
		return
	}

	hotel, er := hotelService.HotelService.GetHotelById(uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hotel": hotel})

}

func InsertHotel(c *gin.Context) {
	var payload dto.Hotel
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotel, er := hotelService.HotelService.InsertHotel(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	SendMessage(hotel.HotelID, "CREATE")

	businessURL := "http://business:8080/api/business/mapping-hotel"

	// JSON body
	body := Bodystruct{
		HotelID: hotel.HotelID,
		AmadeusID: hotel.AmadeusID,
	}

	jsonData, _ := json.Marshal(body)

	// Create a HTTP post request
	r, err := http.NewRequest("POST", businessURL, bytes.NewBuffer(jsonData))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to make request!"})
		}

	c.JSON(http.StatusCreated, gin.H{"hotel": hotel})
}

func UpdateHotel(c *gin.Context) {

	uuid, err := uuid.Parse(c.Param("hotelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "HotelID must be a uuid"})
		return
	}

	var payload dto.Hotel
	errr := c.BindJSON(&payload)
	if errr != nil {
		c.JSON(http.StatusBadRequest, errr.Error())
		return
	}

	payload.HotelID = uuid

	hotel, er := hotelService.HotelService.UpdateHotel(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	SendMessage(hotel.HotelID, "UPDATE")

	c.JSON(http.StatusOK, gin.H{"hotel": hotel})
}

func DeleteHotel(c *gin.Context) {

	uuid, err := uuid.Parse(c.Param("hotelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "HotelID must be a uuid"})
		return
	}

	er := hotelService.HotelService.DeleteHotel(uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	SendMessage(uuid, "DELETE")

	c.JSON(http.StatusOK, gin.H{"success": "Hotel deleted successfully"})

}

func UploadPhoto(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("hotelID"))
	newFileName := uuid.New().String()

	uuid, err := uuid.Parse(c.Param("hotelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotelID must be a uuid"})
		return
	}

	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fileExt := filepath.Ext(file.Filename)

	file.Filename = "/images/hotels/" + newFileName + fileExt
	log.Debug("file name: ", file.Filename)

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, "static"+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var payload dto.Photo
	payload.Url = file.Filename

	photo, er := hotelService.HotelService.InsertPhoto(payload, uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"photo": photo})
}
