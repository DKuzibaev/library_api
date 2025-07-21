package main

import (
	"library_api/internal/server"
	"log"
)

func main() {
	log.Println("Library API Service is starting...")
	// TODO: Конфигурация сервиса.
	server := server.NewLibraryAPI(nil) // Здесь должен быть реальный репозиторий, а не nil.

	// TODO: Запуск сервиса.

	// TODO: Остановка сервиса.
}
