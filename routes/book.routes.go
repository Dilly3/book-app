package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dilly3/book-app/database"
	"github.com/dilly3/book-app/models"
	utils "github.com/dilly3/book-app/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	store database.Datastore
}

func (h *Handler) CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book = new(models.Book)
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error reading request body",
				Message: fmt.Sprint(err),
			})
			return
		}
		err = json.Unmarshal(body, &book)
		if err != nil {
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error unmarshalling request body",
				Message: fmt.Sprint(err),
			})
			return
		}

		newBook, err := h.store.AddBook(book)
		if err != nil {
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error creating book",
				Message: fmt.Sprint(err),
			})
			return
		}
		c.JSON(200, utils.SuccessResponse{
			Code:    http.StatusOK,
			Object:  newBook,
			Message: "Book created successfully",
		})

	}
}

func (h *Handler) GetBook() gin.HandlerFunc {
	return func(c *gin.Context) {

		bookId := c.Param("book_id")

		var book *models.Book

		objectId, _ := primitive.ObjectIDFromHex(bookId)

		book, err := h.store.GetBook(objectId)
		if err != "" {
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

func (h *Handler) UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {

		bookId := c.Param("book_id")
		var book models.Book
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error reading request body",
				Message: fmt.Sprint(err),
			})
			return
		}
		err = json.Unmarshal(body, &book)
		if err != nil {
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error unmarshalling request body",
				Message: fmt.Sprint(err),
			})
			return
		}

		objectId, _ := primitive.ObjectIDFromHex(bookId)
		updatedBook, err := h.store.UpdateBook(objectId, &book)
		if err != nil {
			c.JSON(500, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Error:   "Error updating book",
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
