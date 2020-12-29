package handler

import (
	"github.com/gin-gonic/gin"
)

type BruteHandler struct{}

func (handler BruteHandler) Show(c *gin.Context) {
	c.JSON(200, gin.H{
		"phrase":        "et tu",
		"language_code": "la",
	})
	return
}
