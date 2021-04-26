package main

import (
	bruteHandler "github.com/bryutus/brute/app/handler/brute"
	"github.com/bryutus/brute/app/infrastructure"
	"github.com/bryutus/brute/app/infrastructure/validator"
	"github.com/gin-gonic/gin"
)

func main() {
	db := infrastructure.Init()
	defer db.Close()

	router().Run()
}

func router() *gin.Engine {
	r := gin.Default()

	validator.Register()

	brute := r.Group("/brute")
	{
		showHandler := bruteHandler.ShowHandler{}
		brute.GET("", func(c *gin.Context) {
			showHandler.Show(c)
		})

		createHandler := bruteHandler.CreateHandler{}
		brute.POST("", func(c *gin.Context) {
			createHandler.Create(c)
		})

		updateHandler := bruteHandler.UpdateHandler{}
		brute.PUT("", func(c *gin.Context) {
			updateHandler.Update(c)
		})
	}

	return r
}
