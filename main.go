package main

import (
	"database/sql"
	"log"

	"go-user-api/config"
	"go-user-api/db/sqlc"
	"go-user-api/internal/handler"
	"go-user-api/internal/logger"
	"go-user-api/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err := db.Ping(); err != nil {
	log.Fatal("DB connection failed:", err)
}

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := sqlc.New(db)
	logg := logger.New()

	app := fiber.New()

	userHandler := &handler.UserHandler{
		Queries:  queries,
		Validate: validator.New(),
		Logger:   logg,
	}

	routes.Register(app, userHandler)

	log.Println("Server running on :3000")
	log.Fatal(app.Listen(":3000"))
	app.Get("/", func(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "User API running",
	})
})

}
