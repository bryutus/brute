package handler

import (
	"net/http"

	"github.com/bryutus/brute/app/infrastructure/persistence"
	"github.com/bryutus/brute/app/usecase"
	"github.com/gin-gonic/gin"
)

type BruteHandler struct{}

type requestShowAphorism struct {
	LanguageCode string `form:"language_code" binding:"omitempty,len=2,iso639_1_alpha2,not_exists_language_code"`
}

func (handler BruteHandler) Show(c *gin.Context) {
	var req requestShowAphorism
	if err := c.ShouldBind(&req); err != nil {
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

type requestCreateAphorism struct {
	Phrase       string `form:"phrase" binding:"required,min=1"`
	LanguageCode string `form:"language_code" binding:"required,len=2,iso639_1_alpha2,exists_language_code"`
}

func (handler BruteHandler) Create(c *gin.Context) {
	var req requestCreateAphorism
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

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

type requestUpdateAphorism struct {
	Phrase       string `form:"phrase" binding:"required,min=1"`
	LanguageCode string `form:"language_code" binding:"required,len=2,iso639_1_alpha2,not_exists_language_code"`
}

func (handler BruteHandler) Update(c *gin.Context) {
	var req requestUpdateAphorism
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

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
