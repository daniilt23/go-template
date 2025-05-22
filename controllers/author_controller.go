package controllers

import (
	"net/http"

	"go-test/models"
	"go-test/services"

	"github.com/gin-gonic/gin"
)

func GetAllAuthors(c *gin.Context) {
	var author []models.Author
	err := services.GetAllAuthors(&author)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, author)
}

func GetAuthorById(c *gin.Context) {
	id := c.Param("id")
	var  author models.Author
	err := services.GetAuthorByID(&author, id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, author)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author
	var err error = c.ShouldBindJSON(&author)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = services.CreateAuthor(&author)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, author)
}

func UpdateAuthor(c *gin.Context) {
    id := c.Param("id")
    var author models.Author
    if err := c.ShouldBindJSON(&author); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := services.UpdateAuthor(&author, id)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"name": author.Name})
}

func DeleteAuthor(c *gin.Context) {
    id := c.Param("id")
    var author models.Author
    err := services.DeleteAuthor(&author, id)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.IndentedJSON(http.StatusOK, gin.H{"message": "Author deleted"})
}