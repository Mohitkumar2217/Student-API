package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MohitKumar2217/Students-api/internal/config"
	"github.com/MohitKumar2217/Students-api/internal/http/handlers/student"
	"github.com/MohitKumar2217/Students-api/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// logger
	// database setup
	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initilized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))

	// setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}
	slog.Info("Server started", slog.String("address", cfg.HTTPServer.Address))

	// channel for sync
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// gracefully shutdown using go routine
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	// channel initialied
	<-done

	// server stop logic
	slog.Info("Shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx) // gracefully shutdown method
	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")
}
