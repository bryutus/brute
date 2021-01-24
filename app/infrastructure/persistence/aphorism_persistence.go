package persistence

import (
	"fmt"

	"github.com/bryutus/brute/app/domain/model"
	"github.com/bryutus/brute/app/domain/repository"
	"github.com/bryutus/brute/app/infrastructure"
	"github.com/bryutus/brute/app/infrastructure/dto"
)

type aphorismPersistence struct{}

func NewAphorismPersistence() repository.AphorismRepository {
	return &aphorismPersistence{}
}

func (bp aphorismPersistence) FindBy(code string) (*model.Aphorism, error) {
	db := infrastructure.GetDB()
	aphorismDTO := dto.Aphorism{}

	if result := db.Where("language_code = ?", code).First(&aphorismDTO); result.Error != nil {
		return nil, fmt.Errorf("record not found: language_code=%s", code)
	}

	return aphorismDTO.ConvertToModel(), nil
}

func (bp aphorismPersistence) Save(code string, phrase string) (*model.Aphorism, error) {
	db := infrastructure.GetDB()
	aphorismDTO := dto.Aphorism{}

	result := db.Where("language_code = ?", code).
		Assign(dto.Aphorism{
			LanguageCode: code,
			Phrase:       phrase}).
		FirstOrCreate(&aphorismDTO)

	if result.Error != nil {
		return nil, fmt.Errorf("create error: language_code=%s, phrase=%s", code, phrase)
	}
	return aphorismDTO.ConvertToModel(), nil
}
