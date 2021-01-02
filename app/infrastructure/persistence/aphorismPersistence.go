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
