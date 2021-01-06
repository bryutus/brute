package handler

import (
	"net/http"

	"github.com/bryutus/brute/app/infrastructure/persistence"
	"github.com/bryutus/brute/app/usecase"
	"github.com/gin-gonic/gin"
)

type BruteHandler struct{}

func (handler BruteHandler) Show(code string, c *gin.Context) {
	aphorismRepository := persistence.NewAphorismPersistence()
	usecase := usecase.NewBruteUseCaseImplement(aphorismRepository)

	result, err := usecase.Exec(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"phrase":        result.Phrase,
		"language_code": result.LanguageCode,
	})
	return
}
