package usecase

import (
	"github.com/bryutus/brute/app/domain/model"
	"github.com/gin-gonic/gin"
)

type BruteUseCase interface {
	Find(*gin.Context) (model.Brute, error)
}

type bruteUseCase struct{}

func NewBruteUseCase() BruteUseCase {
	return &bruteUseCase{}
}

func (usecase bruteUseCase) Find(c *gin.Context) (model.Brute, error) {
	brute := &model.Brute{
		Phrase:       "et tu",
		LanguageCode: "la",
	}
	return *brute, nil
}
