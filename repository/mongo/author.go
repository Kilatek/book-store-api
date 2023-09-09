package mongorepo

import (
	"context"
	"log"
	"time"

	entities "bookstore.com/domain/entity"
	portError "bookstore.com/port/error"
	"bookstore.com/repository"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const AuthorCollectionName = "authors"

type authorRepository struct {
	client  *mongo.Client
	db      string
	timeout time.Duration
}

func NewAuthorRepository(mongoServerURL, mongoDb string, timeout int) (repository.AuthorRepository, error) {
	mongoClient, err := newMongClient(mongoServerURL, timeout)
	repo := &authorRepository{
		client:  mongoClient,
		db:      mongoDb,
		timeout: time.Duration(timeout) * time.Second,
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to new author mongo repository")
	}

	return repo, nil
}

func (r *authorRepository) Store(ctx context.Context, author *entities.Author) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	collection := r.client.Database(r.db).Collection(AuthorCollectionName)

	now := time.Now()
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"_id":         primitive.NewObjectID(),
			"firstName":   author.FirstName,
			"lastName":    author.LastName,
			"birthDate":   author.BirthDate,
			"nationality": author.Nationality,
			"createdAt":   now,
			"updatedAt":   now,
		},
	)
	if err != nil {
		return errors.Wrap(err, "authorRepository.Store")
	}

	return nil
}

func (r *authorRepository) Update(ctx context.Context, author *entities.Author) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(author.Id)
	if err != nil {
		return portError.NewBadRequestError("unable to parse author ID to ObjectID", err)
	}

	collection := r.client.Database(r.db).Collection(AuthorCollectionName)
	now := time.Now()
	_, err = collection.UpdateByID(
		ctx,
		_id,
		bson.D{
			{
				Key: "$set", Value: bson.D{
					{Key: "firstName", Value: author.FirstName},
					{Key: "lastName", Value: author.LastName},
					{Key: "birthDate", Value: author.BirthDate},
					{Key: "nationality", Value: author.Nationality},
					{Key: "updatedAt", Value: now},
				},
			},
		},
	)
	if err != nil {
		return errors.Wrap(err, "authorRepository.Update")
	}

	return nil
}

func (r *authorRepository) Find(ctx context.Context, id string) (*entities.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, portError.NewBadRequestError("unable to parse author ID to ObjectID", err)
	}

	author := &entities.Author{}
	collection := r.client.Database(r.db).Collection("authors")

	filter := bson.M{"_id": _id}
	err = collection.FindOne(ctx, filter).Decode(author)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, portError.NewNotFoundError("author not found", err)
		}
		return nil, errors.Wrap(err, "authorRepository.Find")
	}

	return author, nil

}

func (r *authorRepository) FindAll(ctx context.Context) ([]*entities.Author, error) {
	var authors []*entities.Author
	collection := r.client.Database(r.db).Collection(AuthorCollectionName)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, &authors); err != nil {
		return nil, errors.Wrap(err, "authorRepository.FindAll")
	}

	return authors, nil
}

func (r *authorRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return portError.NewBadRequestError("unable to parse author ID to ObjectID", err)
	}

	filter := bson.M{"_id": _id}
	collection := r.client.Database(r.db).Collection(AuthorCollectionName)
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
