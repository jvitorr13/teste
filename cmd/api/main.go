package main

import (
	"log"
	"net/http"

	taskDelivery "taskMetrics/internal/task/delivery"
	taskUC "taskMetrics/internal/task/usecase"
	userDelivery "taskMetrics/internal/user/delivery"
	userUC "taskMetrics/internal/user/usecase"
	"taskMetrics/pkg/config"
	"taskMetrics/pkg/database"
	"taskMetrics/pkg/logger"
)

func main() {
	// Load config
	cfg := config.Load()

	// Initialize logger
	l := logger.New()

	// Initialize database (pool)
	db, err := database.NewConnection(cfg.DBURL)
	if err != nil {
		l.Fatal("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize User Module (usecase com transação por chamada e search_path seguro)
	userUsecase := userUC.NewUserUseCase(db)
	uDelivery := userDelivery.NewUserDelivery(userUsecase)

	// Initialize Task Module (usecase com transação por chamada e search_path seguro)
	taskUsecase := taskUC.NewTaskUseCase(db)
	tDelivery := taskDelivery.NewTaskDelivery(taskUsecase)

	// Set up routes
	mux := http.NewServeMux()

	// User routes
	mux.HandleFunc("POST /users", uDelivery.Create)
	mux.HandleFunc("GET /users", uDelivery.List)

	// Task routes
	mux.HandleFunc("POST /tasks", tDelivery.Create)
	mux.HandleFunc("PUT /tasks/status", tDelivery.UpdateStatus)

	// Start server
	l.Info("Starting server on %s", cfg.ServerAddr)
	if err := http.ListenAndServe(cfg.ServerAddr, mux); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
