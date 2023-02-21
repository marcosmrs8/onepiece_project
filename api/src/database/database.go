package database

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
)

func Initialize(ctx context.Context) (*firestore.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	projectId := os.Getenv("PROJECTID")
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}
	if err != nil {
		panic(err)
	}
	return client, nil
}
