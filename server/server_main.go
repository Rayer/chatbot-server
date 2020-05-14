package main

import (
	apiLib "ChatBotServer"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moogar0880/problems"
	_ "github.com/rayer/chatbot"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	chatbot := apiLib.NewChatbotContext()

	r.NoRoute(func(c *gin.Context) {
		err404 := problems.NewStatusProblem(404)
		err404.Detail = "No such route!"
		c.JSON(404, err404)
	})

	r.NoMethod(func(c *gin.Context) {
		err404 := problems.NewStatusProblem(404)
		err404.Detail = "No such method!"
		c.JSON(404, err404)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.POST("/chatbot", func(c *gin.Context) {
		var conv apiLib.ChatbotConversion
		err := c.BindJSON(&conv)

		utx, _ := chatbot.GetUserContext(conv.User)

		prompt, keywords_v, keywords_iv, err := utx.RenderMessageWithDetail()
		str, err := utx.HandleMessage(conv.Input)
		next, err := utx.RenderMessage()
		c.JSON(http.StatusOK, gin.H{
			"prompt":           prompt,
			"keywords":         keywords_v,
			"invalid_keywords": keywords_iv,
			"message":          str,
			"error":            err,
			"next":             next,
		})
	})

	r.DELETE("/chatbot/:user", func(c *gin.Context) {
		user := c.Param("user")
		chatbot.ExpireUser(user, func() {
			c.JSON(201, gin.H{
				"message": "ok",
			})
		}, func() {
			c.JSON(http.StatusBadRequest, problems.NewDetailedProblem(http.StatusBadRequest, "User "+user+" not found!"))
		})

	})

	r.Run()
}
