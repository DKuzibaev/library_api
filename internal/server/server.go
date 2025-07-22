package server

import (
	"library_api/internal/domain/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetBooksList() ([]models.Book, error)
	SaveBook(book models.Book) // Сохранение книги в репозиторий
	UpdateBook(id string, updatedBook models.Book) error // изменение книги 
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

func (LAPI *LibraryAPI) Run() error {
	log.Printf("Library API Service is running on %s", LAPI.httpServer.Addr)
	err := LAPI.httpServer.ListenAndServe()
	return err
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
		books.POST("/create", LAPI.newBook) // Создание новой книги по ID
		books.GET("/list", LAPI.booksList)  // Получение информации всех книгах
		books.GET("/get/:bookID")           // Получение информации о книге по ID
		books.PUT("/update/:bookID", LAPI.updateBook)        // Обновление информации о книге
		books.DELETE("/delete/:bookID")     // Удаление книги по ID
	}

	LAPI.httpServer.Handler = router
}

func (lAPI *LibraryAPI) booksList(ctx *gin.Context) {
	books, err := lAPI.repo.GetBooksList()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func (lAPI *LibraryAPI) newBook(cxt *gin.Context) {
	var book models.Book
	err := cxt.ShouldBindBodyWithJSON(&book)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lAPI.repo.SaveBook(book)
	cxt.JSON(http.StatusCreated, book)

}

func (lAPI *LibraryAPI) updateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var updated models.Book

	if err := ctx.ShouldBindJSON(&updated); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := lAPI.repo.UpdateBook(bookID, updated)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}
