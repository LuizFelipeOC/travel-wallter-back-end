package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizfelipeeoliveiraac/travel-wallter-back-end/internal/auth"
	"github.com/luizfelipeeoliveiraac/travel-wallter-back-end/internal/expense"
	"github.com/luizfelipeeoliveiraac/travel-wallter-back-end/internal/infrastrucuture/database"
	"github.com/luizfelipeeoliveiraac/travel-wallter-back-end/internal/travel"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		fmt.Println("Erro ao conectar no banco:", err)
		return
	}
	defer db.Close()

	authRepo := auth.NewAuthRepository(db)
	authHandler := auth.NewAuthHandler(authRepo)

	travelRepo := travel.NewTravelRepository(db)
	travelHandler := travel.NewTravelHandler(travelRepo)

	expenseRepo := expense.NewExpenseRepository(db)
	expenseHandler := expense.NewExpenseHandler(expenseRepo)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		authGroup.GET("/profile", auth.AuthMiddleware(), authHandler.GetProfile)
	}

	protectedRoutes := router.Group("")
	protectedRoutes.Use(auth.AuthMiddleware())

	travelGroup := protectedRoutes.Group("/travels")
	{
		travelGroup.POST("", travelHandler.CreateTravel)
		travelGroup.GET("", travelHandler.ListTravels)
		travelGroup.GET("/:id", travelHandler.GetTravel)
		travelGroup.PUT("/:id", travelHandler.UpdateTravel)
		travelGroup.DELETE("/:id", travelHandler.DeleteTravel)
	}

	expenseGroup := protectedRoutes.Group("/expenses")
	{
		expenseGroup.POST("/:travel_id", expenseHandler.CreateExpense)
		expenseGroup.GET("/:travel_id", expenseHandler.ListExpenses)
		expenseGroup.GET("/:travel_id/total", expenseHandler.GetTravelTotal)
		expenseGroup.GET("/:travel_id/:id", expenseHandler.GetExpense)
		expenseGroup.PUT("/:travel_id/:id", expenseHandler.UpdateExpense)
		expenseGroup.DELETE("/:travel_id/:id", expenseHandler.DeleteExpense)
	}

	log.Println("✓ Servidor iniciado em :8080")
	router.Run(":8080")
}