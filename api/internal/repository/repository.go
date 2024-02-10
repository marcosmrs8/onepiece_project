package repository

import (
	"api/internal/domain/entity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CharacterRepository interface {
	GetAllCharacters(ctx context.Context) ([]entity.CharsMongo, error)
	GetCharactersByName(ctx context.Context, name string, options interface{}) ([]entity.CharsMongo, error)
}

type characterRepository struct {
	collection *mongo.Collection
}

func NewCharacterRepository(collection *mongo.Collection) CharacterRepository {
	return &characterRepository{collection: collection}
}

func (r *characterRepository) GetAllCharacters(ctx context.Context) ([]entity.CharsMongo, error) {
	var characters []entity.CharsMongo
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &characters); err != nil {
		return nil, err
	}

	return characters, nil
}

func (r *characterRepository) GetCharactersByName(ctx context.Context, name string, options interface{}) ([]entity.CharsMongo, error) {
	var characters []entity.CharsMongo
	filter := bson.M{"name": name}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &characters); err != nil {
		return nil, err
	}
	cursor.Close(ctx)

	// If no documents found by complete name, try to find by partial name in splitedName
	if len(characters) == 0 {
		filter := bson.M{"splitedName": bson.M{"$regex": name, "$options": "i"}}
		cursor, err := r.collection.Find(ctx, filter)
		if err != nil {
			return nil, err
		}
		if err := cursor.All(ctx, &characters); err != nil {
			return nil, err
		}
		cursor.Close(ctx)
	}

	return characters, nil
}
