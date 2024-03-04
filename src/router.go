package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type sourceText struct {
	Source      string `form:"Source" binding:"required"`
	Target      string `form:"Target" binding:"required"`
	PostContent string `form:"postContent" binding:"required"`
}

func routing() {
	router := gin.Default()
	// router.Use(RequestLogger())
	// router.Use(ResponseLogger())
	router.Use(cors.Default())
	router.POST("/sourceText", postsourceText)

	router.Run("localhost:8080")
}

func postsourceText(c *gin.Context) {
	var newSourceText sourceText

	if err := c.ShouldBind(&newSourceText); err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
		return
	}

	// DO something with request body thats now store in newSourceText
	response := handleTranslation(newSourceText)
	c.IndentedJSON(response.respCode, response.retStr)

}
