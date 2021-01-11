package main

import (
	"github.com/bryutus/brute/app/handler"
	"github.com/bryutus/brute/app/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	db := infrastructure.Init()
	defer db.Close()

	r := gin.Default()

	brute := r.Group("/brute")
	{
		BruteHandler := handler.BruteHandler{}

		brute.GET("", func(c *gin.Context) {
			BruteHandler.Show(c)
		})

		brute.POST("", func(c *gin.Context) {
			BruteHandler.Create(c)
		})
	}

	r.Run()
}
