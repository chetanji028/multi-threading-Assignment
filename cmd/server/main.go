package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chetanji028/distributed-file-storage/internal/handlers"
	"github.com/chetanji028/distributed-file-storage/internal/repository"
	"github.com/chetanji028/distributed-file-storage/internal/service"
	"github.com/chetanji028/distributed-file-storage/internal/utils"
)

func main() {
	// Load environment variables
	loadEnv()

	// Connect to the database
	db := utils.ConnectDB()
	defer db.Close()

	// Initialize repository, service, and handlers
	fileRepo := repository.NewFileRepository(db)
	fileService := service.NewFileService(fileRepo)
	fileHandler := handlers.NewFileHandler(fileService)

	// Set up routes
	http.HandleFunc("/upload", fileHandler.UploadFileHandler)
	http.HandleFunc("/file", fileHandler.GetFileDataHandler)
	http.HandleFunc("/download", fileHandler.DownloadFileHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func loadEnv() {
	// Load environment variables from .env file if needed
	// Alternatively, set them in docker-compose.yml
}
