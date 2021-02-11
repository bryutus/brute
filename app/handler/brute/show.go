package brute

import (
	"net/http"

	"github.com/bryutus/brute/app/infrastructure/persistence"
	"github.com/bryutus/brute/app/usecase"
	"github.com/gin-gonic/gin"
)

type ShowHandler struct{}

type requestShowAphorism struct {
	LanguageCode string `form:"language_code" binding:"omitempty,len=2,iso639_1_alpha2,not_exists_language_code"`
}

func (handler ShowHandler) Show(c *gin.Context) {
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
