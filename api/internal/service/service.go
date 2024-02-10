package services

import (
	"api/internal/domain/entity"
	"api/internal/repository"
	"context"
)

type CharacterService interface {
	GetAllCharacters(ctx context.Context) ([]entity.CharsMongo, error)
	GetCharactersByName(ctx context.Context, name string, options interface{}) ([]entity.CharsMongo, error)
}
type characterService struct {
	characterRepository repository.CharacterRepository
}

func NewCharacterService(characterRepository repository.CharacterRepository) CharacterService {
	return &characterService{characterRepository: characterRepository}
}

func (s *characterService) GetAllCharacters(ctx context.Context) ([]entity.CharsMongo, error) {
	return s.characterRepository.GetAllCharacters(ctx)
}

func (s *characterService) GetCharactersByName(ctx context.Context, name string, options interface{}) ([]entity.CharsMongo, error) {
	return s.characterRepository.GetCharactersByName(ctx, name, options)
}
