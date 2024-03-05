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
	Deepl       bool   `form:"Deepl" binding:"required"`
}

type ResponseObj struct {
	CognateResponse string
	DeeplResponse   string
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

	var responseObj ResponseObj
	var respCode int
	responseObj.CognateResponse = response.retStr

	if (newSourceText.Deepl) && (response.respCode == 200) {
		deeplResponse := deeplHandler(newSourceText)
		responseObj.DeeplResponse = deeplResponse.retStr

		if deeplResponse.respCode == 400 {
			respCode = 400
		} else {
			respCode = 200
		}

	} else {
		responseObj.DeeplResponse = ""
		respCode = 200
	}

	c.JSON(respCode, responseObj)

}
