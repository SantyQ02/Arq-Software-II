package model


type Amenity struct {
	AmenityID  string `bson:"AmenityID"`
	Title      string    `bson:"Title"`
}

type Amenities []Amenity

// func (amenitie *Amenitie) BeforeCreate(scope *gorm.Scope) error {
// 	return scope.SetColumn("AmenitieID", uuid.New())
// }
