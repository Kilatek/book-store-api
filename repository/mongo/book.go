package mongorepo

import (
	"context"
	"time"

	entities "bookstore.com/domain/entity"
	portError "bookstore.com/port/error"
	"bookstore.com/repository"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const BookCollectionName = "books"

type bookRepository struct {
	client  *mongo.Client
	db      string
	timeout time.Duration
}

func NewBookRepository(mongoServerURL, mongoDb string, timeout int) (repository.BookRepository, error) {
	mongoClient, err := newMongClient(mongoServerURL, timeout)
	repo := &bookRepository{
		client:  mongoClient,
		db:      mongoDb,
		timeout: time.Duration(timeout) * time.Second,
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to new book mongo repository")
	}

	return repo, nil
}

func (r *bookRepository) Store(ctx context.Context, book *entities.Book) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	authorId, err := primitive.ObjectIDFromHex(book.AuthorId)
	if err != nil {
		return portError.NewBadRequestError("unable to parse author ID to ObjectID", err)
	}

	collection := r.client.Database(r.db).Collection(BookCollectionName)

	now := time.Now()
	_, err = collection.InsertOne(
		ctx,
		bson.M{
			"_id":             primitive.NewObjectID(),
			"authorId":        authorId,
			"name":            book.Name,
			"description":     book.Description,
			"publicationDate": book.PublicationDate,
			"price":           book.Price,
			"createdAt":       now,
			"updatedAt":       now,
		},
	)
	if err != nil {
		return errors.Wrap(err, "bookRepository.Store")
	}

	return nil
}

func (r *bookRepository) Update(ctx context.Context, book *entities.Book) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(book.Id)
	if err != nil {
		return portError.NewBadRequestError("unable to parse book ID to ObjectID", err)
	}

	authorId, err := primitive.ObjectIDFromHex(book.AuthorId)
	if err != nil {
		return portError.NewBadRequestError("unable to parse author ID to ObjectID", err)
	}

	collection := r.client.Database(r.db).Collection(BookCollectionName)
	now := time.Now()
	_, err = collection.UpdateByID(
		ctx,
		_id,
		bson.D{
			{
				Key: "$set", Value: bson.D{
					{Key: "authorId", Value: authorId},
					{Key: "name", Value: book.Name},
					{Key: "description", Value: book.Description},
					{Key: "publicationDate", Value: book.PublicationDate},
					{Key: "price", Value: book.Price},
					{Key: "updatedAt", Value: now},
				},
			},
		},
	)
	if err != nil {
		return errors.Wrap(err, "bookRepository.Update")
	}

	return nil
}

func (r *bookRepository) Find(ctx context.Context, id string) (*entities.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, portError.NewBadRequestError("unable to parse author ID to ObjectID", err)
	}

	var books []*entities.Book
	collection := r.client.Database(r.db).Collection(BookCollectionName)
	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "authors",
				"localField":   "authorId",
				"foreignField": "_id",
				"as":           "author",
			},
		},
		{
			"$unwind": "$author",
		},
		{
			"$match": bson.M{
				"_id": _id,
			},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Wrap(err, "bookRepository.FindAll")
	}

	err = cursor.All(ctx, &books)
	if err != nil {
		return nil, errors.Wrap(err, "bookRepository.FindAll")
	}

	if len(books) == 0 {
		return nil, portError.NewNotFoundError("book not found", nil)
	}

	return books[0], nil

}

func (r *bookRepository) FindAll(ctx context.Context) ([]*entities.Book, error) {
	var books []*entities.Book
	collection := r.client.Database(r.db).Collection(BookCollectionName)
	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "authors",
				"localField":   "authorId",
				"foreignField": "_id",
				"as":           "author",
			},
		},
		{
			"$unwind": "$author",
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Wrap(err, "bookRepository.FindAll")
	}

	err = cursor.All(ctx, &books)
	if err != nil {
		return nil, errors.Wrap(err, "bookRepository.FindAll")
	}

	return books, nil
}

func (r *bookRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return portError.NewBadRequestError("unable to parse author ID to ObjectID", err)
	}

	filter := bson.M{"_id": _id}
	collection := r.client.Database(r.db).Collection(BookCollectionName)
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
