package router

import (
	"fon/genai"
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
	router.POST("/signup", signUp)
	router.POST("/login", logIn)
	router.GET("/lessons", getAll)
	router.POST("/lessons/:id/ai", askGPT)
	router.GET("/lessons/search/:query", search)
	router.GET("/lessons/search", getAll)

	router.Run("localhost:8080")
}

func search(c *gin.Context) {
	query := c.Param("query")
	c.IndentedJSON(http.StatusOK, persistence.GetBySearch(query))
}

func getAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persistence.GetAll())
}

func signUp(c *gin.Context) {
	var userAttempt persistence.UserSignup
	c.BindJSON(&userAttempt)

	err := persistence.AddUser(userAttempt)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in user creation!"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "User succesfully created!"})
	}
}

func logIn(c *gin.Context) {
	var userAttempt persistence.UserAttempt
	c.BindJSON(&userAttempt)
	user := persistence.GetUser(userAttempt.Email, userAttempt.Password)

	if user.Email != "" {
		c.IndentedJSON(http.StatusOK, user)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User not found!"})
	}
}

func getLesson(c *gin.Context) {
	idS := c.Param("id")
	id, _ := strconv.Atoi(idS)

	c.IndentedJSON(http.StatusOK, persistence.GetLessonByID(id))
}

func askGPT(c *gin.Context) {
	idS := c.Param("id")
	id, _ := strconv.Atoi(idS)

	var promptRequest genai.PromptRequest
	c.BindJSON(&promptRequest)

	result := genai.LessonAsk(promptRequest.Prompt, id)
	c.IndentedJSON(http.StatusOK, result)
}
