package brute

import (
	"net/http"

	"github.com/bryutus/brute/app/infrastructure/persistence"
	"github.com/bryutus/brute/app/usecase"
	"github.com/gin-gonic/gin"
)

type UpdateHandler struct{}

type requestUpdateAphorism struct {
	Phrase       string `form:"phrase" binding:"required,min=1"`
	LanguageCode string `form:"language_code" binding:"required,len=2,iso639_1_alpha2,not_exists_language_code"`
}

func (handler UpdateHandler) Update(c *gin.Context) {
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
