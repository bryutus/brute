package persistence

import (
	"github.com/bryutus/brute/app/domain/model"
	"github.com/bryutus/brute/app/domain/repository"
)

type brutePersistence struct{}

func NewBrutePersistence() repository.BruteRepository {
	return &brutePersistence{}
}

func (bp brutePersistence) FindBy(code string) (*model.Brute, error) {
	brute := &model.Brute{
		Phrase:       "et tu",
		LanguageCode: code,
	}
	return brute, nil
}
