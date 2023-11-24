package hotelController

import (
	"mvc-go/dto"
	hotelService "mvc-go/services/hotel"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"io/ioutil"
	"bytes"
	e "mvc-go/utils/errors"

)

func initTestService() {
	hotelService.HotelService = &hotelService.HotelMockService{}
}

type BodyRes struct {
	Hotel dto.Hotel `json:"hotel"`
}
type BodyResHotels struct {
	Hotels dto.Hotels `json:"hotels"`
}
type BodyResPhoto struct {
	Photo dto.Photo `json:"photo"`
}
type ErrorRes struct {
	Error string `json:"error"`
}

func TestGetHotelById(t *testing.T) {
	initTestService()
	
	hotelID := uuid.New()

	hotelDto := dto.Hotel{
		HotelID: 	 	hotelID,
		AmadeusID:      "0000",
		Title:         	"Test",
		Description:    "Test desciption",
		PricePerDay:    999,
		CityCode:       "City",
		Photos:         nil,
		Amenities: 		nil,
		Active:         true,
		}
	
	hotelMockService := hotelService.HotelService.(*hotelService.HotelMockService)
	hotelMockService.On("GetHotelById", hotelID).Return(hotelDto, nil)

	router := gin.Default()
	router.GET("/test/gethotelbyid/:hotelID", GetHotelById)

	req, _ := http.NewRequest("GET", "/test/gethotelbyid/" + hotelID.String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	bytes, _ := ioutil.ReadAll(resp.Body)

	var body BodyRes
	json.Unmarshal(bytes, &body)

	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, hotelDto, body.Hotel)
}

func TestGetHotelById_InvalidUUID(t *testing.T) {
	initTestService()

	invalidUUID := "invalid-uuid"

	router := gin.Default()
	router.GET("/test/gethotelbyid/:hotelID", GetHotelById)

	req, _ := http.NewRequest("GET", "/test/gethotelbyid/"+invalidUUID, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	var errorResponse ErrorRes
	err := json.Unmarshal(resp.Body.Bytes(), &errorResponse)
	if err != nil {
		t.Fatal("Error decoding JSON response:", err)
	}

	expectedErrorMessage := "HotelID must be a uuid"
	assert.Equal(t, expectedErrorMessage, errorResponse.Error)
}

func TestInsertHotel_Success(t *testing.T) {
	initTestService()

	// Crear un payload de prueba
	hotelPayload := dto.Hotel{
		AmadeusID:   "0000",
		Title:       "Test",
		Description: "Test description",
		PricePerDay: 999,
		CityCode:    "City",
		Photos:      nil,
		Amenities:   nil,
		Active:      true,
	}

	// Mock del servicio de hotel para simular la inserción exitosa
	hotelMockService := hotelService.HotelService.(*hotelService.HotelMockService)
	hotelMockService.On("InsertHotel", hotelPayload).Return(hotelPayload, nil)

	router := gin.Default()
	router.POST("/test/insertHotel", InsertHotel)

	// Crear una solicitud HTTP con el payload del hotel
	reqBody, err := json.Marshal(hotelPayload)
	if err != nil {
		t.Fatal("Error creating request body:", err)
	}

	req, _ := http.NewRequest("POST", "/test/insertHotel", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Verificar la respuesta exitosa
	assert.Equal(t, http.StatusCreated, resp.Code)

	var successResponse BodyRes
	err = json.Unmarshal(resp.Body.Bytes(), &successResponse)
	if err != nil {
		t.Fatal("Error decoding JSON response:", err)
	}

	// Verificar que la respuesta contenga el hotel esperado
	assert.Equal(t, hotelPayload, successResponse.Hotel)
}

func TestInsertHotel_BindJSONError(t *testing.T) {
	initTestService()

	// Crear una solicitud HTTP con un cuerpo JSON inválido
	invalidJSON := []byte(`"invalid": "json"`)
	req, _ := http.NewRequest("POST", "/test/insertHotel", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	router := gin.Default()
	router.POST("/test/insertHotel", InsertHotel)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Verificar que la respuesta contenga un código de estado 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Verificar que el mensaje de error sea el esperado
	assert.Equal(t, resp.Body.String(), "\"json: cannot unmarshal string into Go value of type dto.Hotel\"")
}

func TestGetHotels_Success(t *testing.T) {
	// Inicializa el servicio y el controlador
	initTestService()
	router := gin.Default()
	router.GET("/test/gethotels", GetHotels)

	// Configura el mock del servicio para devolver hoteles de prueba
	hotelMockService := hotelService.HotelService.(*hotelService.HotelMockService)
	hotels := dto.Hotels{
		dto.Hotel{HotelID: uuid.New(), Title: "Hotel 1"},
		dto.Hotel{HotelID: uuid.New(), Title: "Hotel 2"},
		// Agrega más hoteles según sea necesario
	}
	hotelMockService.On("GetHotels").Return(hotels, nil)

	// Realiza la solicitud HTTP GET
	req, _ := http.NewRequest("GET", "/test/gethotels", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Verifica el código de estado y el cuerpo de la respuesta
	assert.Equal(t, http.StatusOK, resp.Code)

	var body BodyResHotels
	err := json.Unmarshal(resp.Body.Bytes(), &body)
	assert.Nil(t, err)
	assert.Equal(t, hotels, body.Hotels)
}

func TestGetHotels_Error(t *testing.T) {
	// Inicializa el servicio y el controlador
	initTestService()
	router := gin.Default()
	router.GET("/test/gethotels", GetHotels)

	// Configura el mock del servicio para devolver un error
	hotelMockService := hotelService.HotelService.(*hotelService.HotelMockService)
	expectedError := e.NewNotFoundApiError("Hotels not found")
	hotelMockService.On("GetHotels").Return(dto.Hotels{}, expectedError)

	// Realiza la solicitud HTTP GET
	req, _ := http.NewRequest("GET", "/test/gethotels", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Verifica el código de estado y el cuerpo de la respuesta de error
	assert.Equal(t, 404, resp.Code)

	var body ErrorRes
	err := json.Unmarshal(resp.Body.Bytes(), &body)
	assert.Nil(t, err)
	assert.Equal(t, "Hotels not found", body.Error)
}

func TestUpdateHotel_Success(t *testing.T) {
	initTestService()

	// Crear un hotel existente para actualizar
	existingHotelID := uuid.New()
	existingHotel := dto.Hotel{
		HotelID:     existingHotelID,
		AmadeusID:   "0000",
		Title:       "Existing Hotel",
		Description: "Existing Description",
		PricePerDay: 888,
		CityCode:    "City",
		Photos:      nil,
		Amenities:   nil,
		Active:      true,
	}

	// Mock del servicio de hotel para simular la obtención del hotel existente
	hotelMockService := hotelService.HotelService.(*hotelService.HotelMockService)
	hotelMockService.On("GetHotelById", existingHotelID).Return(existingHotel, nil)

	// Crear un payload de actualización
	updatePayload := dto.Hotel{
		HotelID:     existingHotelID,
		AmadeusID:   "0000",
		Title:       "Updated Hotel",
		Description: "Updated Description",
		PricePerDay: 999,
		CityCode:    "Updated City",
		Photos:      nil,
		Amenities:   nil,
		Active:      false,
	}

	// Mock del servicio de hotel para simular la actualización exitosa
	hotelMockService.On("UpdateHotel", updatePayload).Return(updatePayload, nil)

	router := gin.Default()
	router.PUT("/test/updatehotel/:hotelID", UpdateHotel)

	// Crear una solicitud HTTP con el payload de actualización
	reqBody, err := json.Marshal(updatePayload)
	if err != nil {
		t.Fatal("Error creating request body:", err)
	}

	req, _ := http.NewRequest("PUT", "/test/updatehotel/"+existingHotelID.String(), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Verificar la respuesta exitosa
	assert.Equal(t, http.StatusOK, resp.Code)

	var successResponse BodyRes
	err = json.Unmarshal(resp.Body.Bytes(), &successResponse)
	if err != nil {
		t.Fatal("Error decoding JSON response:", err)
	}

	// Verificar que la respuesta contenga el hotel actualizado
	assert.Equal(t, updatePayload, successResponse.Hotel)
}

func TestUpdateHotel_InvalidHotelID(t *testing.T) {
	initTestService()

	// Crear una solicitud con un HotelID no válido
	invalidHotelID := "invalid-uuid"
	updatePayload := dto.Hotel{
		HotelID:     uuid.New(), // HotelID válido para evitar errores en la actualización
		AmadeusID:   "0000",
		Title:       "Updated Hotel",
		Description: "Updated Description",
		PricePerDay: 999,
		CityCode:    "Updated City",
		Photos:      nil,
		Amenities:   nil,
		Active:      false,
	}

	router := gin.Default()
	router.PUT("/test/updatehotel/:hotelID", UpdateHotel)

	// Crear una solicitud HTTP con el payload de actualización y un HotelID no válido
	reqBody, err := json.Marshal(updatePayload)
	if err != nil {
		t.Fatal("Error creating request body:", err)
	}

	req, _ := http.NewRequest("PUT", "/test/updatehotel/"+invalidHotelID, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Verificar que la respuesta sea un código de estado BadRequest (400)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	var errorResponse ErrorRes
	err = json.Unmarshal(resp.Body.Bytes(), &errorResponse)
	if err != nil {
		t.Fatal("Error decoding JSON response:", err)
	}

	// Verificar que la respuesta contenga el mensaje de error esperado
	assert.Equal(t, "HotelID must be a uuid", errorResponse.Error)
}

func TestUpdateHotel_InvalidJSON(t *testing.T) {
	initTestService()

	// Crear una solicitud con un cuerpo JSON no válido
	invalidJSON := []byte("invalid-json")
	hotelID := uuid.New()
	router := gin.Default()
	router.PUT("/test/updatehotel/:hotelID", UpdateHotel)

	// Crear una solicitud HTTP con un cuerpo JSON no válido
	req, _ := http.NewRequest("PUT", "/test/updatehotel/"+hotelID.String(), bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Verificar que la respuesta sea un código de estado BadRequest (400)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	assert.Equal(t, "\"invalid character 'i' looking for beginning of value\"",resp.Body.String())
}

func TestDeleteHotel_Success(t *testing.T) {
	initTestService()

	// Crear un UUID para el hotel
	hotelID := uuid.New()

	// Mock del servicio de hotel para simular la eliminación exitosa
	hotelMockService := hotelService.HotelService.(*hotelService.HotelMockService)
	hotelMockService.On("DeleteHotel", hotelID).Return(nil)

	router := gin.Default()
	router.DELETE("/test/deletehotel/:hotelID", DeleteHotel)

	// Crear una solicitud HTTP para eliminar el hotel
	req, _ := http.NewRequest("DELETE", "/test/deletehotel/"+hotelID.String(), nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Verificar la respuesta exitosa
	assert.Equal(t, http.StatusOK, resp.Code)

	var successResponse gin.H
	err := json.Unmarshal(resp.Body.Bytes(), &successResponse)
	if err != nil {
		t.Fatal("Error decoding JSON response:", err)
	}

	// Verificar que la respuesta contenga el mensaje de éxito esperado
	assert.Equal(t, gin.H{"success": "Hotel deleted successfully"}, successResponse)
}

func TestDeleteHotel_Error(t *testing.T) {
	initTestService()

	// Crear un UUID para el hotel
	hotelID := uuid.New()

	// Mock del servicio de hotel para simular un error al eliminar
	expectedError := e.NewInternalServerApiError("Something went wrong deleting hotel", nil)
	hotelMockService := hotelService.HotelService.(*hotelService.HotelMockService)
	hotelMockService.On("DeleteHotel", hotelID).Return(expectedError)

	router := gin.Default()
	router.DELETE("/test/deletehotel/:hotelID", DeleteHotel)

	// Crear una solicitud HTTP para eliminar el hotel
	req, _ := http.NewRequest("DELETE", "/test/deletehotel/"+hotelID.String(), nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Verificar el código de estado y el mensaje de error
	assert.Equal(t, expectedError.Status(), resp.Code)

	var errorResponse ErrorRes
	err := json.Unmarshal(resp.Body.Bytes(), &errorResponse)
	if err != nil {
		t.Fatal("Error decoding JSON response:", err)
	}

	// Verificar que la respuesta contenga el mensaje de error esperado
	assert.Equal(t,expectedError.Message(), errorResponse.Error)
}



