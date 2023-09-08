package entity

import (
	"time"
)

type Author struct {
	Id          string    `json:"id" bson:"_id"`
	FirstName   string    `json:"firstName" bson:"firstName"`
	LastName    string    `json:"lastName" bson:"lastName"`
	BirthDate   string    `json:"birthDate" bson:"birthDate"`
	Nationality string    `json:"nationality" bson:"nationality"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}
