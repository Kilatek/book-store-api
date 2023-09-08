package payload

import (
	"fmt"

	"bookstore.com/tools/datetime"
)

type AuthorRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	BirthDate   string `json:"birthDate"`
	Nationality string `json:"nationality"`
}

func (r *AuthorRequest) Validate() error {
	if r.FirstName == "" {
		return fmt.Errorf("firstName: field required")
	}

	if r.LastName == "" {
		return fmt.Errorf("lastName: field required")
	}

	if r.BirthDate == "" {
		return fmt.Errorf("birthDate: field required")
	}

	_, err := datetime.ParseDate(r.BirthDate)
	if err != nil {
		return fmt.Errorf("birthDate: %s", err)
	}

	if r.Nationality == "" {
		return fmt.Errorf("nationality: field required")
	}

	return nil
}

type AuthorResponse struct {
	Id          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	BirthDate   string `json:"birthDate"`
	Nationality string `json:"nationality"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
