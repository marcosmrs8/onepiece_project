package controllers

import (
	"api/src/entity"
	"api/src/mongodb"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CharsMongo struct {
	Status      string   `bson:"status,omitempty"`
	Name        string   `bson:"name,omitempty"`
	Image       string   `bson:"image,omitempty"`
	Image_manga string   `bson:"image_manga,omitempty"`
	SplitedName []string `bson:"splitedName,omitempty"`
	Bounty      string   `bson:"bounty,omitempty"`
	_id         string   `bson:"_id,omitempty"`
}

func TodosPersonagensMongo(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	mongoClient, err := mongodb.ConnectDB(ctx)
	if err != nil {
		panic(err)
	}
	coll := mongoClient.Collection("one_piece")
	cursor, err := coll.Find(ctx, bson.M{})
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)
	var todos []CharsMongo
	for cursor.Next(ctx) {
		var episode CharsMongo
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, episode)
	}
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		return
	}

}

func PersonageSpecificMongo(w http.ResponseWriter, r *http.Request) {
	var p entity.SpecificChar
	err := json.NewDecoder(r.Body).Decode(&p)
	if p.Name == "" {
		http.Error(w, `faltou o nome`, http.StatusNotFound)
		return
	}
	ctx := context.Background()
	mongoClient, err := mongodb.ConnectDB(ctx)
	if err != nil {
		panic(err)
	}
	coll := mongoClient.Collection("one_piece")
	cursor, err := coll.Find(ctx, bson.M{"name": p.Name})
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal("Erro ao desconectar cursor")
		}
	}(cursor, ctx)
	var episodesFiltered []CharsMongo
	if err = cursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	if episodesFiltered == nil {
		var treatedName = cases.Title(language.Und, cases.NoLower).String(p.Name)
		cursor2, _ := coll.Find(ctx, bson.D{
			{"splitedName", bson.D{{"$all", bson.A{treatedName}}}},
		})
		defer func(cursor2 *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {

			}
		}(cursor2, ctx)
		if err = cursor2.All(ctx, &episodesFiltered); err != nil {
			log.Fatal(err)
		}
		if episodesFiltered == nil {
			http.Error(w, `Nome inválido`, http.StatusNotFound)
			return
		}
	}
	err = json.NewEncoder(w).Encode(episodesFiltered)
	if err != nil {
		return
	}
}
