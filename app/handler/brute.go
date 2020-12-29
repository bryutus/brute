package handler

import (
	"github.com/bryutus/brute/app/usecase"
	"github.com/gin-gonic/gin"
)

type BruteHandler struct{}

func (handler BruteHandler) Show(c *gin.Context) {
	usecase := usecase.NewBruteUseCase()

	result, err := usecase.Find(c)
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
