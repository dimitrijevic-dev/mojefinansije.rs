package router

import (
	"fon/persistence"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
	}))

	router.GET("/lessons/:id", getLesson)

	router.Run("localhost:8080")
}

func getLesson(c *gin.Context) {
	idS := c.Param("id")
	id, _ := strconv.Atoi(idS)

	c.IndentedJSON(http.StatusOK, persistence.GetLessonByID(id))
}
