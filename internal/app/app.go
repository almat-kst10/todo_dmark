package app

import (
	"context"
	"log"
	"todo_dmark/configs"
	"todo_dmark/internal/handler"
	"todo_dmark/internal/repository"
	"todo_dmark/internal/server"
	"todo_dmark/internal/service"
	"todo_dmark/pkg/db"
)

// const timeout = 10 * time.Second

func Run(cfg *configs.Config) error {
	ctx := context.Background()

	postgres, err := db.NewDatabase(ctx, cfg)
	if err != nil {
		log.Println("db.NewDatabase(): ", err)
		return err
	}

	defer postgres.Close()

	// Repo
	todoRepo := repository.NewTodoRepo(postgres)

	// Service
	todoService := service.NewTodoService(todoRepo)

	// Handler
	handler := handler.NewHandler(todoService)
	
	srv := server.NewServer(cfg, handler.Init())
	err = srv.Run()
	if err != nil {
		log.Println("error occurred while running http server: \n", err)
		return err
	}

	return nil
}
