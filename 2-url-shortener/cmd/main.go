package main

import (
	"context"
	"os"
	"sd_concepts/url_shortener/internal/controllers"
	"sd_concepts/url_shortener/internal/repositories"
	"sd_concepts/url_shortener/internal/routes"
	"sd_concepts/url_shortener/internal/services"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	err := godotenv.Load()
	logger.Infof("godotenv.Load() error: %v", err)

	dbURL := os.Getenv("DB_URL")
	logger.Infof("DB_URL from env: %s", dbURL)

	if dbURL == "" {
		logger.Errorf("ERROR: DB_URL is empty!")
		return
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		logger.Errorf("Error creating DB pool - %s", err)
		return
	}
	defer pool.Close()

	repo := repositories.NewRepository(pool)
	svc := services.NewService(repo)
	ctl := controllers.NewController(svc)

	routes.Register(r, ctl)

	if err := r.Run(":3000"); err != nil {
		logger.Errorf("Error running the app - %s", err)
	}
}
