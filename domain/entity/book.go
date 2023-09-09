package entity

import "time"

type Book struct {
	Id              string    `json:"id" bson:"_id"`
	AuthorId        string    `json:"authorId" bson:"authorId"`
	Author          *Author   `json:"author" bson:"author"`
	Name            string    `json:"name" bson:"name"`
	Description     string    `json:"description" bson:"description"`
	PublicationDate string    `json:"publicationDate" bson:"publicationDate"`
	Price           float64   `json:"price" bson:"price"`
	CreatedAt       time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt" bson:"updatedAt"`
}
