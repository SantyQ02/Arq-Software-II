package client

import (
	"fmt"
	"mvc-go/dto"

	solrDB "mvc-go/solr"

	"github.com/google/uuid"
	solr "github.com/rtt/Go-Solr"
	// "net/http"
	// log "github.com/sirupsen/logrus"
)

type solrClient struct{}

type solrClientInterface interface {
	SearchHotels(city string) []dto.Hotel
	UpdateHotel(hotelId string) dto.Hotel
	AddHotel(hotelDto dto.Hotel) dto.Hotel
}

var (
	SolrClient solrClientInterface
)

func init() {
	SolrClient = &solrClient{}
}

func (c *solrClient) SearchHotels(city string) []dto.Hotel {
	query := &solr.Query{
		Params: solr.URLParamMap{
			"q":    []string{fmt.Sprintf("city:\"%s\"", city)}, // Consulta con filtro por ciudad
			"rows": []string{"1000"},                           // Número máximo de filas a recuperar (ajusta según tus necesidades)
		},
	}

	// Realiza la consulta a Solr
	resp, err := solrDB.SolrClient.Select(query)
	if err != nil {
		return []dto.Hotel{}
	}

	// Itera a través de los resultados y construye una lista de hoteles
	var hotels []dto.Hotel
	for _, doc := range resp.Results.Collection {
		hotel := dto.Hotel{
			HotelID:     doc.Fields["hotel_id"].(uuid.UUID),
			AmadeusID:   doc.Field("amadeus_id").(string),
			Title:       doc.Field("title").([]interface{})[0].(string),
			CityCode:    doc.Field("city_code").([]interface{})[0].(string),
			Description: doc.Field("description").([]interface{})[0].(string),
			PricePerDay: doc.Field("price_per_day").(float64),
			Photos:      doc.Field("photos").(dto.Photos),
			Amenities:   doc.Field("amenities").(dto.Amenities),
			Active:      doc.Field("active").(bool),
		}
		hotels = append(hotels, hotel)
	}

	return hotels
}

func (c *solrClient) UpdateHotel(hotelId string) dto.Hotel {
	var hotel dto.Hotel
	return hotel
}

func (c *solrClient) AddHotel(hotelDto dto.Hotel) dto.Hotel {
	var hotel dto.Hotel
	return hotel
}
