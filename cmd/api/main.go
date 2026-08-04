package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luizfelipeeoliveiraac/travel-wallter-back-end/internal/auth"
)

func main() {
	repo := auth.NewAuthRepository()
	handler := auth.NewAuthHandler(repo)

	router := gin.Default()
	
	router.POST("/auth/register", handler.Register)
	router.POST("/auth/login", handler.Login)

	router.Run(":8080")
}