package main

import (
	"beverage-classifier/auto"
	"beverage-classifier/internal"
	"beverage-classifier/internal/db"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = database.AutoMigrate(
		&auto.User{},
		&internal.Beverage{},
		&internal.Ingredient{},
		&internal.Nutrition{},
		&internal.Flag{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	if err := internal.EnsureConstraints(database); err != nil {
		log.Fatal("Constraint bootstrap failed:", err)
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	internal.RegisterHandlers(router, database)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("system ready to work on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
