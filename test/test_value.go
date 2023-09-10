package test

import "time"

const (
	AuthorId1          = "64fbf00fc3a88d3a02b964dc"
	AuthorFirstName1   = "author firstname 1"
	AuthorLastName1    = "author lastname 1"
	AuthorBirthDate1   = "1985-04-04"
	AuthorNationality1 = "Viet Nam"
)

const (
	BookId1          = "8sfbf00fc3a3jd3a02b964ds"
	BookName1        = "book name 1"
	BookDescription1 = "description 1"
	PublicationDate1 = "1992-01-01"
	Price1           = 12.12
	InvalidPrice1    = -1
)

const (
	CreatedAtStr = "2023-09-09T04:09:51.491Z"
	UpdatedAtStr = "2023-09-09T04:09:51.491Z"
)

var (
	CreatedAt, _ = time.Parse(time.RFC3339, CreatedAtStr)
	UpdatedAt, _ = time.Parse(time.RFC3339, UpdatedAtStr)
)
