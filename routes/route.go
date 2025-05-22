package routes

import (
	"github.com/gin-gonic/gin"
	"go-test/controllers"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	versionRouter := router.Group("/api/v1")
	bookRouter(versionRouter)
	authorRouter(versionRouter)

	return router
}

func bookRouter(router *gin.RouterGroup) {
	router.GET("/getAllBooks", controllers.GetBooks)
    router.GET("/getBookById/:id", controllers.GetBookById)
    router.POST("/createBook", controllers.CreateBook)
    router.PUT("/updateBookById/:id", controllers.UpdateBook)
    router.DELETE("/deleteBookById/:id", controllers.DeleteBook)
}

func authorRouter(router *gin.RouterGroup) {
	router.GET("/getAllAuthors", controllers.GetAllAuthors)
    router.GET("/getAuthorById/:id", controllers.GetAuthorById)
    router.POST("/createAuthor", controllers.CreateAuthor)
    router.PUT("/updateAuthorById/:id", controllers.UpdateAuthor)
    router.DELETE("/deleteAuthorById/:id", controllers.DeleteAuthor)
} 