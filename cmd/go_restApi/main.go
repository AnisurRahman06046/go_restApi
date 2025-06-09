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

	"github.com/AnisurRahman06046/go_restApi/internal/config"
	"github.com/AnisurRahman06046/go_restApi/internal/http/handlers/student"
	"github.com/AnisurRahman06046/go_restApi/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("storage initialized", slog.String("env", cfg.Env))

	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/student/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))
	router.HandleFunc("PATCH /api/student/{id}", student.UpdateById(storage))
	router.HandleFunc("DELETE /api/student/{id}", student.DeleteById(storage))

	server := &http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	log.Printf("Server is running at %s\n", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start server")
		}
	}()
	<-done

	slog.Info("shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shut down", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")

}
