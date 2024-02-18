package server

import (
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config *viper.Viper
	router *chi.Mux
}

func InitHttpServer(config *viper.Viper) *HttpServer {
	return &HttpServer{
		config: config,
		router: chi.NewRouter(),
	}
}

func (s *HttpServer) Start() {
	s.router.Use(middleware.Logger)

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	err := http.ListenAndServe(":8080", s.router)

	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
		return
	}
}
