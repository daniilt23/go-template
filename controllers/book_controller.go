package controllers

import (
	"net/http"

	"go-test/models"
	"go-test/services"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	err := services.GetAllBooks(&books)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	var  book models.Book
	err := services.GetBookByID(&book, id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	var err error = c.ShouldBindJSON(&book)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = services.CreateBook(&book)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := services.UpdateBook(&book, id)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    err := services.DeleteBook(&book, id)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted"})
}