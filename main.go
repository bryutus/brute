package main

import (
	"github.com/bryutus/brute/app/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	brute := r.Group("/brute")
	{
		BruteHandler := handler.BruteHandler{}
		brute.GET("", BruteHandler.Show)
	}

	r.Run()
}
