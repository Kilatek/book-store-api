package payload

import (
	"fmt"

	"bookstore.com/tools/datetime"
)

type BookRequest struct {
	AuthorId        string  `json:"authorId"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	PublicationDate string  `json:"publicationDate"`
	Price           float64 `json:"price"`
}

func (r *BookRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name: field required")
	}

	if r.Description == "" {
		return fmt.Errorf("description: field required")
	}

	if r.PublicationDate == "" {
		return fmt.Errorf("publicationDate: field required")
	}

	_, err := datetime.ParseDate(r.PublicationDate)
	if err != nil {
		return fmt.Errorf("publicationDate: %s", err)
	}

	if r.Price == 0 {
		return fmt.Errorf("price: field required")
	}

	if r.Price < 0 {
		return fmt.Errorf("price: invalid")
	}

	if r.AuthorId == "" {
		return fmt.Errorf("authorId: field required")
	}

	return nil
}

type BookResponse struct {
	Id              string          `json:"id"`
	Author          *AuthorResponse `json:"author"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	PublicationDate string          `json:"publicationDate"`
	Price           float64         `json:"price"`
	CreatedAt       string          `json:"createdAt"`
	UpdatedAt       string          `json:"updatedAt"`
}
