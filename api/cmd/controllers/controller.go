package controllers

import (
	services "api/internal/service"
	"context"
	"encoding/json"
	"net/http"
)

type CharacterController struct {
	characterService services.CharacterService
}

func NewCharacterController(characterService services.CharacterService) *CharacterController {
	return &CharacterController{characterService: characterService}
}

func (c *CharacterController) GetAllCharacters(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	characters, err := c.characterService.GetAllCharacters(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(characters); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *CharacterController) GetCharacterByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	character, err := c.characterService.GetCharactersByName(ctx, name, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if character == nil {
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(character); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
