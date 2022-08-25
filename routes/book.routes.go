package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/dilly3/book-app/database"
	"github.com/dilly3/book-app/models"
	utils "github.com/dilly3/book-app/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type Handle struct {
	store  database.DataStore
	Logger *zap.Logger
}

func NewHandle(databaseFactory func() database.DataStore) *Handle {
	db := databaseFactory()
	return &Handle{
		store:  db,
		Logger: zap.NewExample(),
	}
}
func (h *Handle) CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book = new(models.Book)
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Println(err)
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error reading request body",
				Message: fmt.Sprint(err),
			})
			return
		}
		err = json.Unmarshal(body, &book)
		if err != nil {
			log.Println(err)
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error unmarshalling request body",
				Message: fmt.Sprint(err),
			})
			return
		}

		newBook, err := h.store.AddBook(book)
		if err != nil {
			log.Println(err)
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error creating book",
				Message: fmt.Sprint(err),
			})
			return
		}
		c.JSON(200, utils.SuccessResponse{
			Code:    http.StatusCreated,
			Object:  newBook,
			Message: "Book created successfully",
		})

	}
}

func (h *Handle) GetBook() gin.HandlerFunc {
	return func(c *gin.Context) {

		bookId := c.Param("book_id")

		var book *models.Book

		objectId, _ := utils.GetPrimitiveObjectId(bookId)

		book, err := h.store.GetBook(objectId)
		if err != nil {
			h.Logger.Info("Error getting book", zap.Error(err.(error)))
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error getting book",
				Message: fmt.Sprint(err),
			})
			return
		}
		c.JSON(200, utils.SuccessResponse{
			Code:    http.StatusOK,
			Object:  book,
			Message: "Book retrieved successfully",
		})
	}

}

func (h *Handle) UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {

		bookId := c.Param("book_id")
		var book models.Book
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Println(err)
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error reading request body",
				Message: fmt.Sprint(err),
			})
			return
		}
		err = json.Unmarshal(body, &book)
		if err != nil {
			h.Logger.Info("Error unmarshalling request body", zap.Error(err))
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error unmarshalling request body",
				Message: fmt.Sprint(err),
			})
			return
		}

		objectId, _ := primitive.ObjectIDFromHex(bookId)
		_, err = h.store.UpdateBook(objectId, &book)
		if err != nil {
			h.Logger.Info("Error updating book", zap.Error(err))
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error updating book",
				Message: fmt.Sprint(err),
			})
			return
		}
		updatedBook, errStr := h.store.GetBook(objectId)
		if errStr != nil {
			h.Logger.Info("Error getting book", zap.Error(errStr))
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error getting book",
				Message: fmt.Sprint(err),
			})
			return
		}

		c.JSON(200, utils.SuccessResponse{
			Code:    http.StatusOK,
			Object:  updatedBook,
			Message: "Book updated successfully",
		})
	}
}

func (h *Handle) GetAllBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := h.store.GetAllBooks()
		if err != nil {
			h.Logger.Info("Error getting all books", zap.Error(err))
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error getting all books",
				Message: fmt.Sprint(err),
			})
			return
		}
		c.JSON(200, utils.SuccessResponse{
			Code:    http.StatusOK,
			Object:  books,
			Message: "Books retrieved successfully",
		})

	}
}

func (h *Handle) DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {

		bookId := c.Param("_id")

		objectId, _ := utils.GetPrimitiveObjectId(bookId)
		book, err := h.store.GetBook(objectId)
		if err != nil {
			h.Logger.Info("Error getting book", zap.Error(err))
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error getting book",
				Message: fmt.Sprint(err),
			})
			return
		}
		err2 := h.store.DeleteBook(objectId)
		if err2 != nil {
			h.Logger.Info("Error deleting book", zap.Error(err))
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error deleting book",
				Message: fmt.Sprint(err),
			})
			return
		}
		c.JSON(200, utils.SuccessResponse{
			Code:    http.StatusOK,
			Message: fmt.Sprintf("%s deleted successfully", *book.Title),
		})

	}
}
