// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/AnisurRahman06046/go_restApi/internal/config"
// )

// func main() {
// 	cfg := config.MustLoad()

// 	router := http.NewServeMux()
// 	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Welcome to home page"))
// 	})

// 	// server setup
// 	server := http.Server{
// 		Addr:    cfg.Addr,
// 		Handler: router,
// 	}
// 	fmt.Println("Server is running")
// 	err := server.ListenAndServe()
// 	if err != nil {
// 		log.Fatalf("failed to start server")
// 	}

// }

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
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New())

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

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shut down", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")

}
