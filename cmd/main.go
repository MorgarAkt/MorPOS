package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MorgarAkt/MorPOS/internal/handlers"
	"github.com/MorgarAkt/MorPOS/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	services.InitDB("./sqlite.db")
	defer services.CloseDB()

	userService := services.NewUserService(services.Db)
	userHandler := handlers.NewUserHandler(userService)

	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not found on .env file!")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlers.HandlerReadiness)
	v1Router.Get("/err", handlers.HandlerErr)
	v1Router.Get("/err", handlers.HandlerErr)
	v1Router.Post("/users", userHandler.SaveUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Server starting on port: %s\n", port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
