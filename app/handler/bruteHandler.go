package handler

import (
	"net/http"

	"github.com/bryutus/brute/app/domain/model"

	"github.com/bryutus/brute/app/infrastructure/persistence"
	"github.com/bryutus/brute/app/usecase"
	"github.com/gin-gonic/gin"
)

type BruteHandler struct{}

func (handler BruteHandler) Show(c *gin.Context) {
	var aphorism model.Aphorism
	if err := c.ShouldBind(&aphorism); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	code := c.DefaultQuery("language_code", "la")

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

func (handler BruteHandler) Create(c *gin.Context) {
	code := c.PostForm("language_code")
	phrase := c.PostForm("phrase")

	aphorismRepository := persistence.NewAphorismPersistence()
	usecase := usecase.NewSaveBruteUseCaseImplement(aphorismRepository)

	restult, err := usecase.Exec(code, phrase)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"phrase":        restult.Phrase,
		"language_code": restult.LanguageCode,
	})
	return
}
