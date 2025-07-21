package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Repository interface {
}

type LibraryAPI struct {
	httpServer *http.Server
	repo       Repository
}

func NewLibraryAPI(repo Repository) *LibraryAPI {
	httpServer := http.Server{
		Addr: ":8080", // TODO: заменить на конфигурацию из файла
	}

	lAPI := &LibraryAPI{
		httpServer: &httpServer,
		repo:       repo,
	}

	lAPI.configRoutes()

	return lAPI
}

func (LAPI *LibraryAPI) configRoutes() {
	router := gin.Default()
	users := router.Group("/users")
	{
		users.POST("/register")         // Регистрация нового пользователя
		users.POST("/login")            // Авторизация пользователя
		users.PUT("/update/:userID")    // Обновление информации о пользователе
		users.DELETE("/delete/:userID") // Удаление пользователя
		users.GET("/:userID")           // Получение информации о пользователе
	}

	books := router.Group("/books")
	{
		books.POST("/create")           // Создание новой книги по ID
		books.GET("/get/:bookID")       // Получение информации о книге
		books.PUT("/update/:bookID")    // Обновление информации о книге
		books.DELETE("/delete/:bookID") // Удаление книги по ID
	}

	LAPI.httpServer.Handler = router
}
