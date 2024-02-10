package main

import (
	"api/cmd/routes"
	"api/internal/repository"
	"api/internal/repository/mongodb"
	services "api/internal/service"
	"context"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Iniciando servidor...")

	ctx := context.Background()
	database, err := mongodb.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
	}

	characterRepository := repository.NewCharacterRepository(database.Collection("one_piece"))

	characterService := services.NewCharacterService(characterRepository)

	routes.HandleRequest(characterService)
}
