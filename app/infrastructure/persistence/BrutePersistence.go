package persistence

import (
	"fmt"

	"github.com/bryutus/brute/app/domain/model"
	"github.com/bryutus/brute/app/domain/repository"
	"github.com/bryutus/brute/app/infrastructure"
	"github.com/bryutus/brute/app/infrastructure/dto"
)

type brutePersistence struct{}

func NewBrutePersistence() repository.BruteRepository {
	return &brutePersistence{}
}

func (bp brutePersistence) FindBy(code string) (*model.Brute, error) {
	db := infrastructure.GetDB()
	aphorismDTO := dto.Aphorism{}

	if result := db.Where("language_code = ?", code).First(&aphorismDTO); result.Error != nil {
		return nil, fmt.Errorf("record not found: language_code=%s", code)
	}

	return aphorismDTO.ConvertToModel(), nil
}
