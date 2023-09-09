package entity

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Id        string    `json:"id" bson:"_id"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	FirstName string    `json:"firstName" bson:"firstName"`
	LastName  string    `json:"lastName" bson:"lastName"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
