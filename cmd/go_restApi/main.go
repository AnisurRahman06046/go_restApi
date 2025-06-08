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
	"log"
	"net/http"

	"github.com/AnisurRahman06046/go_restApi/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("Welcome to home page"))
	})

	server := &http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	log.Printf("Server is running at %s\n", cfg.HTTPServer.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
