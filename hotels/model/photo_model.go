package model


type Photo struct {
	PhotoID string `bson:"PhotoID"`
	Url     string    `bson:"Url"`
}

type Photos []Photo

// func (photo *Photo) BeforeCreate(scope *gorm.Scope) error {
// 	return scope.SetColumn("PhotoID", uuid.New())
// }
