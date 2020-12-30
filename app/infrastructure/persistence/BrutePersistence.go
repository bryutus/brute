package persistence

import (
	"github.com/bryutus/brute/app/domain/model"
	"github.com/bryutus/brute/app/domain/repository"
)

type brutePersistence struct{}

func NewBrutePersistence() repository.BruteRepository {
	return &brutePersistence{}
}

func (bp brutePersistence) FindBy() (*model.Brute, error) {
	brute := &model.Brute{
		Phrase:       "et tu",
		LanguageCode: "la",
	}
	return brute, nil
}
