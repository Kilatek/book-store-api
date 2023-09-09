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

const UserCollectionName = "users"

type userRepository struct {
	client  *mongo.Client
	db      string
	timeout time.Duration
}

func NewUserRepository(mongoServerURL, mongoDb string, timeout int) (repository.UserRepository, error) {
	mongoClient, err := newMongClient(mongoServerURL, timeout)
	repo := &userRepository{
		client:  mongoClient,
		db:      mongoDb,
		timeout: time.Duration(timeout) * time.Second,
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to new author mongo repository")
	}

	return repo, nil
}

func (r *userRepository) Store(ctx context.Context, user *entities.User) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	collection := r.client.Database(r.db).Collection(UserCollectionName)

	now := time.Now()
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"_id":       primitive.NewObjectID(),
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"password":  user.Password,
			"username":  user.Username,
			"createdAt": now,
			"updatedAt": now,
		},
	)
	if err != nil {
		return errors.Wrap(err, "mongoRepository.Store")
	}

	return nil
}

func (r *userRepository) Find(ctx context.Context, username string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	user := &entities.User{}
	collection := r.client.Database(r.db).Collection(UserCollectionName)

	filter := bson.M{"username": username}
	err := collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, portError.NewNotFoundError("user not found", err)
		}
		return nil, errors.Wrap(err, "mongoRepository.Find")
	}

	return user, nil

}
