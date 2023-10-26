package client

import (
	"fmt"
	"errors"
	"mvc-go/dto"

	solrDB "mvc-go/solr"

	"github.com/google/uuid"
	solr "github.com/rtt/Go-Solr"
	// "net/http"
	log "github.com/sirupsen/logrus"
	"encoding/json"
)

type solrClient struct{}

type solrClientInterface interface {
	SearchHotels(city string) ([]dto.Hotel, error)
	AddOrUpdateHotel(hotelDto dto.Hotel) error
	EmptyCollection() error
}

var (
	SolrClient solrClientInterface
)

func init() {
	SolrClient = &solrClient{}
}

func (c *solrClient) SearchHotels(city string) ([]dto.Hotel, error) {
	query := &solr.Query{
		Params: solr.URLParamMap{
			"q":    []string{fmt.Sprintf("city_code:\"%s\"", city)}, // Consulta con filtro por ciudad
		},
	}

	// Realiza la consulta a Solr
	resp, err := solrDB.SolrClient.Select(query)
	if resp == nil {
		return []dto.Hotel{}, nil
	}
	if err != nil {
		return []dto.Hotel{}, err
	}

	// Itera a través de los resultados y construye una lista de hoteles
	var hotels []dto.Hotel
	for _, doc := range resp.Results.Collection {
		id, err := uuid.Parse(doc.Fields["id"].(string))
		if err != nil {
			return []dto.Hotel{}, errors.New("Invalid UUID")
		}
		hotel := dto.Hotel{
			HotelID:     id,
			AmadeusID:   doc.Field("amadeus_id").([]interface{})[0].(string),
			Title:       doc.Field("title").([]interface{})[0].(string),
			CityCode:    doc.Field("city_code").([]interface{})[0].(string),
			Description: doc.Field("description").([]interface{})[0].(string),
			Thumbnail:   getThumbnailFromInterface(doc.Field("thumbnail")),
			PricePerDay: doc.Field("price_per_day").([]interface{})[0].(float64),
			Photos:      getPhotosFromInterface(doc.Field("photos")),
			Amenities:   getAmenitiesFromInterface(doc.Field("amenities")),
			Active:      doc.Field("active").([]interface{})[0].(bool),
		}
		hotels = append(hotels, hotel)
	}
	if hotels == nil {
		return dto.Hotels{}, nil
	}

	return hotels, nil
}

func (c *solrClient) EmptyCollection() error {

	f := map[string]interface{}{
		"delete": map[string]interface{}{
			"query":    "*:*",
		},
	}

	// Realiza la consulta a Solr
	_, err := solrDB.SolrClient.Update(f, true)
	if err != nil {
		return err
	}

	return nil
}

func (c *solrClient) AddOrUpdateHotel(hotelDto dto.Hotel) error {
    var amenitiesStr, photosStr *string
    var thumbnail string

    if hotelDto.Photos != nil && len(hotelDto.Photos) > 0 {
        if photos, err := json.Marshal(hotelDto.Photos); err == nil {
            s := string(photos)
            photosStr = &s
            thumbnail = hotelDto.Photos[0].Url
        }
    }

    if hotelDto.Amenities != nil {
        if amenities, err := json.Marshal(hotelDto.Amenities); err == nil {
            s := string(amenities)
            amenitiesStr = &s
        }
    }

	hotelDocument := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{
				"id":          hotelDto.HotelID,
				"amadeus_id": hotelDto.AmadeusID,
				"title":        hotelDto.Title,
				"city_code":        hotelDto.CityCode,
				"description": hotelDto.Description,
				"price_per_day": hotelDto.PricePerDay,
				"amenities": amenitiesStr,
				"photos": photosStr,
				"thumbnail":    thumbnail,
				"active": hotelDto.Active,
			},
		},
	}

	// Inserta el nuevo documento en Solr
	_, err := solrDB.SolrClient.Update(hotelDocument, true) // El segundo parámetro "true" realiza una confirmación inmediata
	if err != nil {
		return err
	}
	return nil
}

func getPhotosFromInterface(i interface{}) dto.Photos {
	if i == nil {
		return dto.Photos{}
	}
	if photos, ok := i.([]interface{}); ok {
		if len(photos) > 0 {
			if a, ok := photos[0].(string); ok {
				var dto dto.Photos
				json.Unmarshal([]byte(a), &dto)
				return dto
			}
		}
	}
	log.Error("Failed to convert interface to dto.Photos")
	return dto.Photos{}
}

func getAmenitiesFromInterface(i interface{}) dto.Amenities {
	if i == nil {
		return dto.Amenities{}
	}
	if amenities, ok := i.([]interface{}); ok {
		if len(amenities) > 0 {
			if a, ok := amenities[0].(string); ok {
				var dto dto.Amenities
				json.Unmarshal([]byte(a), &dto)
				return dto
			}
		}
	}
	log.Error("Failed to convert interface to dto.Amenities")
	return dto.Amenities{}
}

func getThumbnailFromInterface(i interface{}) string {
	if i == nil {
        return ""
    }
    if slice, ok := i.([]interface{}); ok {
        if len(slice) > 0 {
            if str, isString := slice[0].(string); isString {
                return str
            }
        }
    }
    return ""
}