package persistence

import (
	"errors"
	"fmt"

	"github.com/bryutus/brute/app/domain/model"
	"github.com/bryutus/brute/app/domain/repository"
	"github.com/bryutus/brute/app/infrastructure"
	"github.com/bryutus/brute/app/infrastructure/dto"
	"github.com/jinzhu/gorm"
)

type aphorismPersistence struct{}

func NewAphorismPersistence() repository.AphorismRepository {
	return &aphorismPersistence{}
}

func (bp aphorismPersistence) FindBy(code string) (*model.Aphorism, error) {
	db := infrastructure.GetDB()
	aphorismDTO := dto.Aphorism{}

	res := db.Where("language_code = ?", code).First(&aphorismDTO)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if res.Error != nil {
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
