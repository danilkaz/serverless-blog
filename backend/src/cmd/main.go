package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"serverless-blog/internal/config"
	"serverless-blog/internal/db/ydb"
	"serverless-blog/internal/handlers"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalf("Config load error: %v", err)
	}

	ydbClient, err := ydb.New(config.YDB.Endpoint, config.YDB.ServiceAccountKey)
	if err != nil {
		log.Fatalf("YDB client creation error: %v", err)
	}

	ctx := context.Background()

	if err := ydbClient.CreateTable(ctx); err != nil {
		log.Fatalf("Table creation error: %v", err)
	}

	hostname, _ := os.Hostname()

	versionHandler := handlers.NewVersionHandler(config.Version, hostname+"-"+uuid.NewString())
	createPostHandler := handlers.NewCreatePostHandler(ydbClient)
	listPostsHandler := handlers.NewListPostsHandler(ydbClient)

	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	router.HandleFunc("/version", versionHandler.Handle)
	router.HandleFunc("/post", createPostHandler.Handle).Methods("POST")
	router.HandleFunc("/post", listPostsHandler.Handle)

	corsConfig := cors.AllowAll()

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: corsConfig.Handler(router),
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Server shutdown error: %v", err)
		}
	}()

	log.Printf("Server started on port %s", config.Port)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Server listen and serve error: %v", err)
	}

	log.Print("Bye!")
}
