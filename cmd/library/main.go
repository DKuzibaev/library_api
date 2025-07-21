package main

import (
	"library_api/internal/repository/inmemory"
	"library_api/internal/server"
	"log"
)

func main() {
	log.Println("Library API Service is starting...")
	// TODO: Конфигурация сервиса.
	repo := inmemory.NewLocalStorage()   // Инициализация репозитория
	server := server.NewLibraryAPI(repo) // Здесь должен быть реальный репозиторий, а не nil.

	// TODO: Запуск сервиса.
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run Library API Service: %v", err)
	}

	// TODO: Остановка сервиса.
	log.Println("Library API Service has stopped...")
}
