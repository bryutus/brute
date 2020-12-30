package handler

import (
	"github.com/bryutus/brute/app/infrastructure/persistence"
	"github.com/bryutus/brute/app/usecase"
	"github.com/gin-gonic/gin"
)

type BruteHandler struct{}

func (handler BruteHandler) Show(code string, c *gin.Context) {
	bruteRepository := persistence.NewBrutePersistence()
	usecase := usecase.NewBruteUseCaseImplement(bruteRepository)

	result, err := usecase.Exec(code)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"phrase":        result.Phrase,
		"language_code": result.LanguageCode,
	})
	return
}
