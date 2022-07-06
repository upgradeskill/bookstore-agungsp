package bookcontroller

import (
	"bookstore/controllers/authcontroller"
	"bookstore/models"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type BookController struct {
	sync.Mutex
	model models.BookModels
}

func (bc *BookController) Route(echo *echo.Echo) {
	/**
	 * Middleware setup per controller
	 */
	groupRoute := echo.Group("/books")
	config := middleware.JWTConfig{
		Claims:     &authcontroller.AuthController{},
		SigningKey: []byte("secret"),
	}
	groupRoute.Use(middleware.JWTWithConfig(config))

	/**
	 * Declare the routes
	 */
	echo.GET("/books", bc.index)
	echo.GET("/books/:id", bc.show)
	groupRoute.POST("", bc.create)
	groupRoute.PUT("/:id", bc.update)
	groupRoute.DELETE("/:id", bc.delete)
}

func (bc *BookController) Model(db *gorm.DB) {
	bc.model.Db = db
	bc.model.Init()
}

func (bc *BookController) index(context echo.Context) error {
	books := bc.model.GetAll()
	return context.JSON(http.StatusOK, books)
}

func (bc *BookController) create(context echo.Context) error {
	book := models.Book{
		Isbn: randomString(15),
		Title: context.FormValue("title"),
		Author: context.FormValue("author"),
		Publisher: context.FormValue("publisher"),
		Category: context.FormValue("category"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	bc.model.Create(&book)
	return context.String(http.StatusOK, "Data has been created")
}

func (bc *BookController) show(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.Echo().Logger.Fatal(err)
	}
	book := bc.model.Get(id)

	if book.Id == 0 {
		return context.String(http.StatusNotFound, "Data Not Found!")	
	}

	return context.JSON(http.StatusOK, book)
}

func (bc *BookController) update(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.Echo().Logger.Fatal(err)
	}
	book := bc.model.Get(id)

	if book.Id == 0 {
		return context.String(http.StatusNotFound, "Data Not Found!")	
	}

	if context.FormValue("title") != "" {
		book.Title = context.FormValue("title") 		
	}

	if context.FormValue("author") != "" {
		book.Author = context.FormValue("author") 	
	}

	if context.FormValue("publisher") != "" {
		book.Publisher = context.FormValue("publisher") 
	}

	if context.FormValue("category") != "" {
		book.Category = context.FormValue("category") 
	}

	bc.model.Update(book)

	return context.String(http.StatusOK, "Data has been updated")
}

func (bc *BookController) delete(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.Echo().Logger.Fatal(err)
	}
	book := bc.model.Get(id)

	if book.Id == 0 {
		return context.String(http.StatusNotFound, "Data Not Found!")	
	}

	bc.model.Delete(book)

	return context.String(http.StatusOK, "Data has been deleted")
}

func randomString(n int) string {
    var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
 
    s := make([]rune, n)
    for i := range s {
        s[i] = letters[rand.Intn(len(letters))]
    }
    return string(s)
}