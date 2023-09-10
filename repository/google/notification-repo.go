package google

import (
	"context"
	"fmt"
	"os"

	"bookstore.com/domain/entity"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

type FireDB struct {
	*db.Client
}

var fireDB FireDB

func (db *FireDB) Connect() error {
	// Find home directory.
	home, err := os.Getwd()
	if err != nil {
		return err
	}
	ctx := context.Background()
	opt := option.WithCredentialsFile(home + "/config/firebase-credential.json")
	config := &firebase.Config{DatabaseURL: "https://book-store-5b397-default-rtdb.firebaseio.com/"}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}
	db.Client = client
	return nil
}

func (db *FireDB) Store(ctx context.Context, book *entity.Book) {
	db.NewRef("books/"+book.Id).Set(context.Background(), book)
}

func FirebaseDB() *FireDB {
	return &fireDB
}
